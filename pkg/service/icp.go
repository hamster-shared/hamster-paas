package service

import (
	"database/sql"
	"errors"
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/db"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils"
	"hamster-paas/pkg/utils/logger"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var (
	GetVersion     = "dfx -V"
	AccountId      = "dfx ledger account-id"
	GetPrincipal   = "dfx identity get-principal"
	GetWallet      = "dfx identity get-wallet --network %s"
	NewIdentity    = "dfx identity new %s --storage-mode plaintext"
	UseIdentity    = "dfx identity use %s"
	DeployWallet   = "dfx identity deploy-wallet %s --network %s"
	LedgerBalance  = "dfx ledger balance --network %s" // icp
	WalletBalance  = "dfx wallet balance --network %s" // cycle
	CreateCanister = "dfx ledger create-canister %s --amount %s --network %s "
	WalletTopUp    = "dfx ledger top-up %s --amount %s --network %s"
	DepositCycles  = "dfx canister deposit-cycles %s %s --network %s"
	CanisterStatus = "dfx canister status %s --network %s"

	TransferICP = "dfx ledger transfer %s --icp %s --memo %s --network %s"
)

type IcpService struct {
	db      *gorm.DB
	network string // ic
}

func NewIcpService() *IcpService {
	alineDb, err := application.GetBean[*gorm.DB]("alineDb")
	if err != nil {
		return nil
	}
	return &IcpService{
		db:      alineDb,
		network: "ic",
	}
}

func (i *IcpService) GetDfxVersion() (string, error) {
	return i.execDfxCommand(GetVersion)
}

// api brief
func (i *IcpService) GetAccountBrief(userId uint) (*vo.AccountBrief, error) {
	var projects []db.Project
	err := i.db.Model(db.Project{}).Where("user_id = ?", userId).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	var brief vo.AccountBrief
	var canisters []db.IcpCanister
	for _, proj := range projects {
		err := i.db.Model(db.IcpCanister{}).Where("project_id = ?", proj.Id).Find(&canisters).Error
		if err != nil {
			return nil, err
		}
		brief.Canisters += len(canisters)
		for _, can := range canisters {
			if can.Status == db.Running {
				brief.Running += 1
			}
			if can.Status == db.Stopped {
				brief.Stopped += 1
			}
		}
	}
	return &brief, nil
}

// api overview
func (i *IcpService) GetAccountOverview(userId uint) (*vo.AccountOverview, error) {
	var projects []db.Project
	err := i.db.Model(db.Project{}).Where("user_id = ?", userId).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	var ov vo.AccountOverview
	ov.Projects = len(projects)
	var canisters []db.IcpCanister
	for _, proj := range projects {
		err = i.db.Model(db.IcpCanister{}).Where("project_id = ?", proj.Id).Find(&canisters).Error
		if err != nil {
			return nil, err
		}
		ov.Canisters += len(canisters)
	}
	var userIcp db.UserIcp
	//
	err = i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return nil, err
	}
	// lock k8s
	lock, err := utils.Lock()
	if err != nil {
		return nil, err
	}
	defer utils.Unlock(lock)
	// use identity
	err = i.useIndentity(userIcp.IdentityName)
	if err != nil {
		return nil, err
	}
	// icp and cycle balance
	icps, err := i.icpBalanceWithUnit()
	if err != nil {
		return nil, err
	}
	cycles, err := i.cycleBalanceWithUnit()
	if err != nil {
		return nil, err
	}
	ov.Icps = icps
	ov.Cycles = cycles
	return &ov, nil
}

// api canister page
func (i *IcpService) GetCanisterPage(userId uint, page int, size int) (*vo.UserCanisterPage, error) {
	var canistersPage vo.UserCanisterPage
	var data []vo.UserCanisterVo

	var projects []db.Project
	if err := i.dbUserProjects(userId, &projects); err != nil {
		return nil, err
	}

	var allCanisters []db.IcpCanister
	var canisters []db.IcpCanister
	var canisterProj map[string]string = make(map[string]string)
	for _, proj := range projects {
		if err := i.dbProjCanisters(proj.Id.String(), &canisters); err != nil {
			return nil, err
		}
		allCanisters = append(allCanisters, canisters...)
		// set map of canisters to project name
		for _, canister := range canisters {
			canisterProj[canister.CanisterId] = proj.Name
		}
	}
	canistersPage.Total = len(allCanisters)
	st := (page - 1) * size
	end := st + size
	if end > len(allCanisters) {
		end = len(allCanisters)
	}
	var item vo.UserCanisterVo
	for _, canister := range allCanisters[st:end] {
		item.CanisterId = canister.CanisterId
		item.CanisterName = canister.CanisterName
		item.Cycles = canister.Cycles.String
		item.Status = canister.Status.String()
		item.Project = canisterProj[canister.CanisterId]
		item.UpdateAt = canister.UpdateTime.Time.Format("yyyy/MM/dd HH:mm:ss")
		data = append(data, item)
	}
	canistersPage.Data = data
	canistersPage.Page = page
	canistersPage.PageSize = size

	return &canistersPage, nil
}

func (i *IcpService) newIndentity(identityName string) (err error) {
	newIdentitySprintf := NewIdentity
	newIdentityCmd := fmt.Sprintf(newIdentitySprintf, identityName)
	_, err = i.execDfxCommand(newIdentityCmd)
	return err
}

func (i *IcpService) useIndentity(identityName string) (err error) {
	useIdentitySprintf := UseIdentity
	useIdentityCmd := fmt.Sprintf(useIdentitySprintf, identityName)
	_, err = i.execDfxCommand(useIdentityCmd)
	return err
}

func (i *IcpService) dbUserProjects(userId uint, projects *[]db.Project) error {
	return i.db.Model(db.Project{}).Where("user_id = ?", userId).Find(&projects).Error
}

func (i *IcpService) dbProjCanisters(projId string, canisters *[]db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("project_id = ?", projId).Find(&canisters).Error
}

func (i *IcpService) dbCanisterInfo(canisterId string, canister *db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("canister_id = ?", canisterId).First(&canister).Error
}

// return accountId, principal
func (i *IcpService) getLedgerInfo(identityName string) (string, string, error) {
	lock, err := utils.Lock()
	if err != nil {
		return "", "", err
	}
	defer utils.Unlock(lock)

	i.useIndentity(identityName)
	if err != nil {
		return "", "", err
	}
	accountIdCmd := AccountId
	accountId, err := i.execDfxCommand(accountIdCmd)
	if err != nil {
		return "", "", err
	}
	principalCmd := GetPrincipal
	principal, err := i.execDfxCommand(principalCmd)
	if err != nil {
		return "", "", err
	}

	return accountId, principal, nil
}

func (i *IcpService) getWalletId(identityName string) (walletId string, err error) {
	getWalletSprintf := GetWallet
	getWalletCmd := fmt.Sprintf(getWalletSprintf, i.network)
	output, err := i.execDfxCommand(getWalletCmd)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`([a-z0-9-]+-[a-z0-9-]+-[a-z0-9-]+-[a-z0-9-]+-[a-z0-9-]+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1], nil
	} else {
		return "", errors.New("fail to get walletId")
	}
}

func (i *IcpService) icpBalanceWithUnit() (string, error) {
	ledgerBalanceSprintf := LedgerBalance
	ledgerBalanceCmd := fmt.Sprintf(ledgerBalanceSprintf, i.network)
	balance, err := i.execDfxCommand(ledgerBalanceCmd)
	return strings.TrimSpace(balance), err
}

func (i *IcpService) cycleBalanceWithUnit() (string, error) {
	walletBalanceSprintf := WalletBalance
	walletBalanceCmd := fmt.Sprintf(walletBalanceSprintf, i.network)
	balance, err := i.execDfxCommand(walletBalanceCmd)
	return strings.TrimSpace(balance), err
}

func (i *IcpService) getIcpBalance() (string, error) {
	balanceSprintf := LedgerBalance
	balanceCmd := fmt.Sprintf(balanceSprintf, i.network)
	balance, err := i.execDfxCommand(balanceCmd)
	if err != nil {
		return "", err
	}
	balanceSplit := strings.Split(balance, " ")
	if len(balanceSplit) > 0 {
		amount, err := strconv.ParseFloat(balanceSplit[0], 64)
		if err != nil {
			return "", err
		}
		if amount > 0.0002 {
			amount -= 0.0002
		} else {
			return "", errors.New("insufficient icp balance")
		}
		return strconv.FormatFloat(amount, 'f', 8, 64), nil
	} else {
		return "", errors.New("failure to obtain ICP balances")
	}
}

// deprecated
func (i *IcpService) getICPs() (string, error) {
	balanceSprintf := LedgerBalance
	balanceCmd := fmt.Sprintf(balanceSprintf, i.network)
	balance, err := i.execDfxCommand(balanceCmd)
	if err != nil {
		return "", err
	}
	balanceSplit := strings.Split(balance, " ")
	if len(balanceSplit) > 0 {
		return balanceSplit[0], nil
	} else {
		return "", errors.New("failure to obtain cycle balances")
	}
}

func (i *IcpService) getCycles() (string, error) {
	walletBalanceSprintf := WalletBalance
	walletBalanceCmd := fmt.Sprintf(walletBalanceSprintf, i.network)
	balance, err := i.execDfxCommand(walletBalanceCmd)
	if err != nil {
		return "", err
	}
	balanceSplit := strings.Split(balance, " ")
	if len(balanceSplit) > 0 {
		return balanceSplit[0], nil
	} else {
		return "", errors.New("failure to obtain cycle balances")
	}
}

func (i *IcpService) walletTopUp(identityName string, walletId string) (error error) {
	lock, err := utils.Lock()
	if err != nil {
		return err
	}
	defer utils.Unlock(lock)
	err = i.useIndentity(identityName)
	if err != nil {
		return err
	}
	// TODO all balance?
	balance, err := i.getIcpBalance()
	if err != nil {
		return err
	}
	walletTopUpSprintf := WalletTopUp
	walletTopUpCmd := fmt.Sprintf(walletTopUpSprintf, walletId, balance, i.network)
	output, err := i.execDfxCommand(walletTopUpCmd)
	if err != nil {
		return err
	}
	logger.Infof("identityName-> %s walletId-> %s top-up result is: %s \n", identityName, walletId, output)
	return nil
}

func (i *IcpService) depositCanister(identityName string, cycles string, canisterId string) error {
	lock, err := utils.Lock()
	if err != nil {
		return err
	}
	defer utils.Unlock(lock)
	err = i.useIndentity(identityName)
	if err != nil {
		return err
	}
	depositCyclesSprintf := DepositCycles
	depositCyclesCmd := fmt.Sprintf(depositCyclesSprintf, cycles, canisterId, i.network)
	output, err := i.execDfxCommand(depositCyclesCmd)
	if err != nil {
		return err
	}
	logger.Infof("userid-> %s canisterId-> %s deposit-cycles result is: %s \n", identityName, canisterId, output)
	return nil
}

func (i *IcpService) execDfxCommand(cmd string) (string, error) {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if exitError, ok := err.(*exec.ExitError); ok {
		logger.Errorf("%s Exit status: %d, Exit str: %s", cmd, exitError.ExitCode(), string(exitError.Stderr))
		return "", errors.New(string(exitError.Stderr))
	} else if err != nil {
		// 输出其他类型的错误
		logger.Errorf("%s Failed to execute command: %s", cmd, err)
		return "", err
	}
	logger.Infof("%s Exit result: %s", cmd, string(output))
	return string(output), nil
}

//  Old version

// return if account or wallet id is exist
func (i *IcpService) GetAccountFlag(userId uint) (vo vo.IcpAccountVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	vo.UserId = int(userId)
	vo.WalletIdFlag = false
	vo.AccountIdFlag = false
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vo, nil
		} else {
			return vo, err
		}
	}
	if userIcp.AccountId != "" {
		vo.AccountIdFlag = true
	}
	if userIcp.WalletId != "" {
		vo.WalletIdFlag = true
	}
	return vo, nil
}

// create identity and insert usericp db
func (i *IcpService) CreateIdentity(userId uint) (vo vo.UserIcpInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err == nil {
		return vo, errors.New("you have created an identity")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return vo, err
	}
	// newIdentity
	identityName := strconv.Itoa(int(userId))
	err = i.newIndentity(identityName)
	if err != nil {
		return vo, err
	}
	// getLedger
	aId, pId, err := i.getLedgerInfo(identityName)
	if err != nil {
		return vo, err
	}
	// insert userIcp
	userIcp.FkUserId = userId
	userIcp.IdentityName = identityName
	userIcp.AccountId = strings.TrimSpace(aId)
	userIcp.PrincipalId = strings.TrimSpace(pId)
	userIcp.CreateTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	userIcp.UpdateTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = i.db.Model(db.UserIcp{}).Create(&userIcp).Error
	if err != nil {
		return vo, err
	}
	// result
	vo.UserId = int(userId)
	vo.AccountId = aId
	vo.IcpBalance = "0.0000000 ICP"
	return vo, nil
}

// return accountId and icp balance (ICP)
func (i *IcpService) GetAccountInfo(userId uint) (vo vo.UserIcpInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	// lock k8s
	lock, err := utils.Lock()
	if err != nil {
		return vo, err
	}
	defer utils.Unlock(lock)
	err = i.useIndentity(userIcp.IdentityName)
	if err != nil {
		return vo, err
	}
	balance, err := i.icpBalanceWithUnit()
	if err != nil {
		return vo, err
	}
	// result
	vo.UserId = int(userIcp.FkUserId)
	vo.AccountId = userIcp.AccountId
	vo.IcpBalance = balance
	return vo, nil
}

// return walletId and cycle balance (TC)
func (i *IcpService) GetWalletInfo(userId uint) (vo vo.UserCycleInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	// test
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" && userIcp.WalletId == "icp-test-wallet-id" {
		vo.UserId = int(userIcp.FkUserId)
		vo.CanisterId = userIcp.WalletId
		vo.CyclesBalance = "0.0000000 TC (trillion cycles)"
		return vo, nil
	}
	// lock
	lock, err := utils.Lock()
	if err != nil {
		return vo, err
	}
	defer utils.Unlock(lock)
	err = i.useIndentity(userIcp.IdentityName)
	if err != nil {
		return vo, err
	}
	// get wallet balance
	walletBalanceSprintf := WalletBalance
	walletBalanceCmd := fmt.Sprintf(walletBalanceSprintf, i.network)
	balance, err := i.execDfxCommand(walletBalanceCmd)
	if err != nil {
		return vo, err
	}
	// result
	vo.UserId = int(userIcp.FkUserId)
	vo.CanisterId = userIcp.WalletId
	vo.CyclesBalance = balance
	return vo, nil
}

func (i *IcpService) RechargeWallet(userId uint) (vo vo.UserCycleInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	if userIcp.WalletId == "" {
		walletId, err := i.InitWallet(userIcp)
		if err != nil {
			return vo, err
		}
		userIcp.WalletId = walletId
		err = i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).Updates(&userIcp).Error
		if err != nil {
			return vo, err
		}
	} else {
		err := i.walletTopUp(userIcp.IdentityName, userIcp.WalletId)
		if err != nil {
			return vo, err
		}
	}
	return i.GetWalletInfo(userId)
}

// init account wallet
func (i *IcpService) InitWallet(userIcp db.UserIcp) (walletId string, error error) {
	lock, err := utils.Lock()
	if err != nil {
		return "", err
	}
	defer utils.Unlock(lock)
	err = i.useIndentity(userIcp.IdentityName)
	if err != nil {
		return "", err
	}
	// TODO all balance?
	balance, err := i.getIcpBalance()
	if err != nil {
		return "", err
	}
	// create new canister
	createCanisterSprintf := CreateCanister
	createCanisterCmd := fmt.Sprintf(createCanisterSprintf, userIcp.PrincipalId, balance, i.network)
	output, err := i.execDfxCommand(createCanisterCmd)
	logger.Infof("userid-> %s create-canister result is: %s \n", userIcp.IdentityName, output)
	if err != nil {
		return "", err
	}
	// if no wallet, use canister as wallet, if has wallet, get wallet id
	re := regexp.MustCompile(`Canister created with id: "(.*?)"`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		walletId = matches[1]
	} else {
		walletId, err = i.getWalletId(userIcp.IdentityName)
		if err != nil {
			return "", err
		}
	}
	// deploy wallet
	deployWalletSprintf := DeployWallet
	deployWalletCmd := fmt.Sprintf(deployWalletSprintf, walletId, i.network)
	output, err = i.execDfxCommand(deployWalletCmd)
	logger.Infof("userid-> %s walletId-> %s deploy-wallet result is: %s \n", userIcp.IdentityName, walletId, output)
	if err != nil {
		return "", err
	}
	return walletId, nil
}

func (i *IcpService) QueryIcpCanister(projectId string) (string, error) {
	var data db.IcpCanister
	err := i.db.Model(db.IcpCanister{}).Where("project_id = ?", projectId).First(&data).Error
	if err != nil {
		return "", err
	}
	return data.CanisterId, nil
}

func (i *IcpService) QueryIcpCanisterList(projectId string, page, size int) (*vo.IcpCanisterPage, error) {
	var total int64
	var pageData vo.IcpCanisterPage
	var canisters []db.IcpCanister
	var vo []vo.IcpCanisterVo
	err := i.db.Model(db.IcpCanister{}).Where("project_id = ?", projectId).Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&canisters).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		return &pageData, err
	}
	for _, canister := range canisters {
		logger.Info(canister.Cycles)
		if !canister.Cycles.Valid {
			data, err := i.queryCanisterStatus(canister.CanisterId)
			logger.Debugf("balance data is %s:", data.Balance)
			if err == nil {
				logger.Info("start save balance")
				canister.Cycles = sql.NullString{
					String: data.Balance,
					Valid:  true,
				}
				canister.UpdateTime = sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				}
				i.db.Save(&canister)
			}
		} else {
			isThreeHoursAgo := isTimeThreeHoursAgo(canister.UpdateTime.Time, time.Now())
			if isThreeHoursAgo {
				data, err := i.queryCanisterStatus(canister.CanisterId)
				if err == nil {
					canister.Cycles = sql.NullString{
						String: data.Balance,
						Valid:  true,
					}
					canister.UpdateTime = sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					}
					i.db.Save(&canister)
				}
			}
		}
	}
	copier.Copy(&vo, &canisters)
	pageData.Data = vo
	pageData.Page = page
	pageData.PageSize = size
	pageData.Total = int(total)
	return &pageData, nil
}

func (i *IcpService) queryCanisterStatus(canisterId string) (vo.CanisterStatusRes, error) {
	var res vo.CanisterStatusRes
	canisterStatusSprintf := CanisterStatus
	canisterCmd := fmt.Sprintf(canisterStatusSprintf, canisterId, i.network)
	logger.Infof("exec cmd is %s", canisterCmd)
	cmd := exec.Command("bash", "-c", canisterCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("cmd exec failed: %s", err)
		return res, err
	}
	logger.Infof("status is:%s", string(out))
	re := regexp.MustCompile(`Balance: ([0-9_]+) Cycles`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) > 1 {
		value := matches[1]
		value = strings.ReplaceAll(value, "_", "")
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Errorf("balance parse int failed:%s", err)
			return res, err
		}
		data := float64(number) / math.Pow(10, 12)
		balance := fmt.Sprintf("%.2f\n", data)
		res.Balance = balance
	} else {
		logger.Info("balance not found!")
	}
	statusRegex := regexp.MustCompile(`Status: (.+)`)
	statusMatch := statusRegex.FindStringSubmatch(string(out))
	if len(statusMatch) > 1 {
		res.Status = statusMatch[1]
	} else {
		logger.Info("status not found!")
	}
	return res, nil
}

func isTimeThreeHoursAgo(t time.Time, now time.Time) bool {
	duration := now.Sub(t)
	return duration >= 3*time.Hour
}

func (i *IcpService) SaveDfxJsonData(projectId string, jsonData string) error {
	var dfxData db.IcpDfxData
	err := i.db.Where("project_id = ?", projectId).First(&dfxData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			dfxData.ProjectId = projectId
			dfxData.DfxData = jsonData
			dfxData.CreateTime = sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			err = i.db.Create(&dfxData).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

func (i *IcpService) QueryDfxJsonDataByProjectId(projectId string) (vo.IcpDfxDataVo, error) {
	var data db.IcpDfxData
	var vo vo.IcpDfxDataVo
	err := i.db.Model(db.IcpDfxData{}).Where("project_id = ?", projectId).First(&data).Error
	if err != nil {
		return vo, err
	}
	copier.Copy(&vo, &data)
	return vo, nil
}

func (i *IcpService) IsConfigJsonData(projectId string) bool {
	var data db.IcpDfxData
	err := i.db.Model(db.IcpDfxData{}).Where("project_id = ?", projectId).First(&data).Error
	if err != nil {
		return false
	}
	return true
}

func (i *IcpService) UpdateDfxJsonData(id int, jsonData string) error {
	var data db.IcpDfxData
	err := i.db.Model(db.IcpDfxData{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return err
	}
	data.DfxData = jsonData
	err = i.db.Save(&data).Error
	return err
}
