package service

import (
	"database/sql"
	"errors"
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/db"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
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

// api/icp/account/brief
func (i *IcpService) GetAccountBrief(userId uint) (*vo.AccountBrief, error) {
	var res vo.AccountBrief
	var canisters []db.IcpCanister
	if err := i.dbUserCanisters(userId, &canisters); err != nil {
		return nil, err
	}

	res.Canisters += len(canisters)
	for _, can := range canisters {
		if can.Status == db.Running {
			res.Running += 1
		}
		if can.Status == db.Stopped {
			res.Stopped += 1
		}
	}

	return &res, nil
}

const (
	DEFAULT = "default"
	//MYCANISTER = "t35f6-tiaaa-aaaai-acewq-cai"
)

// api/icp/account/overview
func (i *IcpService) GetAccountOverview(userId uint) (*vo.AccountOverview, error) {
	var res vo.AccountOverview
	var projects []db.Project
	if err := i.dbUserProjects(userId, &projects); err != nil {
		return nil, err
	}
	res.Projects = len(projects)

	var canisters []db.IcpCanister
	if err := i.dbUserCanisters(userId, &canisters); err != nil {
		return nil, err
	}
	res.Canisters += len(canisters)

	// get icp identity
	identityName, err := i.dbIdentityName(userId)
	if err != nil {
		return nil, err
	}

	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
	}

	// icp balance
	icps, err := i.getIcp(identityName)
	if err != nil {
		return nil, err
	}
	res.Icps = icps

	// cycle balance
	cycles, err := i.getCycle(identityName)
	if err != nil {
		return nil, err
	}
	res.Cycles = cycles

	return &res, nil
}

// api/icp/account/canisters
func (i *IcpService) GetCanisterPage(userId uint, page int, size int) (*vo.UserCanisterPage, error) {
	var res vo.UserCanisterPage
	var projects []db.Project
	if err := i.dbUserProjects(userId, &projects); err != nil {
		return nil, err
	}
	var projectName map[string]string = make(map[string]string)
	for _, project := range projects {
		// set map of canisters to project name
		projectName[project.Id.String()] = project.Name

	}
	var canisters []db.IcpCanister
	if err := i.dbUserCanisters(userId, &canisters); err != nil {
		return nil, err
	}

	st := (page - 1) * size
	if st > len(canisters) {
		return nil, errors.New("page out of range")
	}
	end := st + size
	if end > len(canisters) {
		end = len(canisters)
	}

	var data []vo.UserCanisterVo
	var item vo.UserCanisterVo
	for _, canister := range canisters[st:end] {
		item.CanisterId = canister.CanisterId
		item.CanisterName = canister.CanisterName
		item.Cycles = canister.Cycles.String
		item.Status = canister.Status.String()
		item.Project = projectName[canister.ProjectId]
		item.UpdateAt = canister.UpdateTime.Time.Format("2006-01-02 15:04:05")
		data = append(data, item)
	}
	res.Total = len(canisters)
	res.Data = data
	res.Page = page
	res.PageSize = size

	return &res, nil
}

// api/icp/canisters/overview
func (i *IcpService) GetCanisterOverview(userId uint, canisterId string) (*vo.CanisterOverview, error) {
	var res vo.CanisterOverview
	var canister db.IcpCanister
	if err := i.dbCanisterInfo(canisterId, &canister); err != nil {
		return nil, err
	}
	res.CanisterId = canisterId
	res.CanisterName = canister.CanisterName

	var project db.Project
	if err := i.dbProjectInfo(canister.ProjectId, &project); err != nil {
		return nil, err
	}
	res.Project = project.Name

	// get icp identity
	identityName, err := i.dbIdentityName(userId)
	if err != nil {
		return nil, err
	}

	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		canisterId = os.Getenv("CANISTER")
		logger.Debugf("test canisterId: %s", canisterId)
	}

	status, err := i.getCanisterStatus(identityName, canisterId)
	if err != nil {
		return nil, err
	}
	res.Status = status.Status
	res.Cycles = status.Balance
	res.MemorySize = status.MemorySize
	res.ModuleHash = status.ModuleHash
	res.UpdateAt = canister.UpdateTime.Time.Format("2006-01-02 15:04:05")

	return &res, nil
}

// api/icp/account/add-canister
func (i *IcpService) GetContollerPage(userId uint, canisterId string, page int, size int) (*vo.ControllerPage, error) {
	var res vo.ControllerPage
	var userIcp db.UserIcp
	if err := i.dbUserIdentity(userId, &userIcp); err != nil {
		return nil, err
	}
	identityName := userIcp.IdentityName
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		canisterId = os.Getenv("CANISTER")
		logger.Debugf("test canisterId: %s", canisterId)
	}
	status, err := i.getCanisterStatus(identityName, canisterId)
	if err != nil {
		return nil, err
	}
	controllers := status.Controllers

	st := (page - 1) * size
	if st > len(controllers) {
		return nil, errors.New("page out of range")
	}
	end := st + size
	if end > len(controllers) {
		end = len(controllers)
	}

	var data []vo.ControllerVo
	var item vo.ControllerVo
	for _, c := range controllers[st:end] {
		item.PrincipalId = c
		if userIcp.PrincipalId == c {
			item.Type = "Hamster-Managed"
		} else {
			item.Type = "User-Managed"
		}
		data = append(data, item)
	}

	res.Total = len(controllers)
	res.Data = data
	res.Page = page
	res.PageSize = size
	return &res, nil
}

// api/icp/account/add-canister
func (i *IcpService) AddCanister(userId uint, param vo.CreateCanisterParam) error {
	var userIcp db.UserIcp
	if err := i.dbUserIdentity(userId, &userIcp); err != nil {
		return err
	}
	identityName := userIcp.IdentityName
	principalId := userIcp.PrincipalId
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		principalId, _, _ = i.getLedgerInfo(identityName)
		logger.Debugf("test principal: %s", principalId)
	}
	canisterId, err := i.createCanister(identityName, principalId)
	if err != nil {
		return err
	}
	logger.Infof("CREATE canister: %s", canisterId)
	// 添加用户的 canister
	if err := i.dbCreateCanister(userId, canisterId, param.CanisterName); err != nil {
		return err
	}
	return nil
}

// api/icp/account/del-canister
func (i *IcpService) DelCanister(userId uint, param vo.DeleteCanisterParam) error {
	identityName, err := i.dbIdentityName(userId) //获取用户的身份
	if err != nil {
		return err
	}
	canisterId := param.CanisterId
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		canisterId = os.Getenv("CANISTER")
		logger.Debugf("test canisterId: %s", canisterId)
	}
	if err := i.deleteCanister(identityName, canisterId); err != nil {
		return err
	}
	logger.Infof("DELETE canister: %s", canisterId)

	// 删除用户的 canister
	if err := i.dbDeleteCanister(userId, canisterId); err != nil {
		return err
	}
	return nil
}

// api/icp/canister/add-controller
func (i *IcpService) AddController(userId uint, param vo.AddControllerParam) error {
	identityName, err := i.dbIdentityName(userId) //获取用户的身份
	if err != nil {
		return err
	}
	canisterId := param.CanisterId
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		canisterId = os.Getenv("CANISTER")
		logger.Debugf("test canisterId: %s", canisterId)
	}

	if err := i.addController(identityName, canisterId, param.Controller); err != nil {
		return err
	}
	logger.Infof("ADD controller: %s", canisterId)

	return nil
}

// api/icp/canister/del-controller
func (i *IcpService) DelController(userId uint, param vo.DelControllerParam) error {
	identityName, err := i.dbIdentityName(userId) //获取用户的身份
	if err != nil {
		return err
	}
	canisterId := param.CanisterId
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		canisterId = os.Getenv("CANISTER")
		logger.Debugf("test canisterId: %s", canisterId)
	}
	if err := i.delController(identityName, canisterId, param.Controller); err != nil {
		return err
	}
	logger.Infof("DEL controller: %s", canisterId)

	return nil
}

// api/icp/canister/change-status 1. running 2. stop
func (i *IcpService) ChangeCanisterStatus(userId uint, param vo.ChangeStatusParam) error {
	identityName, err := i.dbIdentityName(userId) //获取用户的身份
	if err != nil {
		return err
	}
	canisterId := param.CanisterId
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" {
		identityName = DEFAULT
		canisterId = os.Getenv("CANISTER")
		logger.Debugf("test canisterId: %s", canisterId)
	}
	return i.changeCanisterStatus(identityName, canisterId, param.Status)
}

// TODO
func (i *IcpService) InstallDapp(userId uint, canister vo.InstallDappParam) error {
	return nil
}

// TODO No comsumption yet
func (i *IcpService) GetConsumptionPage(canisterId string, page int, size int) (*vo.ConsumptionPage, error) {
	var cspPage vo.ConsumptionPage
	return &cspPage, nil
}

//	Old version
//
// api/icp/account/get-account
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

// api/icp/account/create-identity
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

// api/icp/account/get-account-info
// return accountId and icp balance (ICP)
func (i *IcpService) GetAccountInfo(userId uint) (vo vo.UserIcpInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	balance, err := i.getIcp(userIcp.IdentityName)
	if err != nil {
		return vo, err
	}
	// result
	vo.UserId = int(userIcp.FkUserId)
	vo.AccountId = userIcp.AccountId
	vo.IcpBalance = balance
	return vo, nil
}

// api/icp/account/get-cycle-info
// return walletId and cycle balance (TC)
func (i *IcpService) GetWalletInfo(userId uint) (vo vo.UserCycleInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	// test
	icpTest := os.Getenv("ICP_TEST")
	if icpTest == "true" && userIcp.WalletId == "icp-test-wallet-id" {
		vo.UserId = int(userIcp.FkUserId)
		vo.CanisterId = userIcp.WalletId
		vo.CyclesBalance = "0.0000000 TC (trillion cycles)"
		return vo, nil
	}
	// get wallet balance
	walletBalanceSprintf := WalletBalance
	walletBalanceCmd := fmt.Sprintf(walletBalanceSprintf, i.network, userIcp.IdentityName)
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

// api/icp/account/buy-cycles
func (i *IcpService) RechargeWallet(userId uint) (vo vo.UserCycleInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	if userIcp.WalletId == "" {
		walletId, err := i.initWallet(userIcp)
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

// api/icp/canister/add-cycles
func (i *IcpService) RechargeCanister(userId uint, param vo.RechargeCanisterParam) (vo vo.UserCycleInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	amount, err := strconv.ParseFloat(param.Amount, 64)
	if err != nil {
		return vo, err
	}
	depositCycles := amount * 1e12
	err = i.depositCanister(userIcp.IdentityName, strconv.FormatFloat(depositCycles, 'f', -1, 64), param.CanisterId)
	if err != nil {
		return vo, err
	}

	data, err := i.getCanisterStatus(userIcp.IdentityName, param.CanisterId)
	if err != nil {
		return vo, err
	}

	var icpCanister db.IcpCanister
	err = i.db.Model(db.IcpCanister{}).Where("canister_id = ?", param.CanisterId).First(&icpCanister).Error
	if err != nil {
		return vo, err
	}
	icpCanister.Cycles = sql.NullString{
		String: data.Balance,
		Valid:  true,
	}
	err = i.db.Model(db.IcpCanister{}).Where("canister_id = ?", param.CanisterId).Updates(&icpCanister).Error
	if err != nil {
		return vo, err
	}
	vo.UserId = int(userId)
	vo.CanisterId = param.CanisterId
	vo.CyclesBalance = data.Balance + "T"
	return vo, nil
}

// use in recharge
// init account wallet
func (i *IcpService) initWallet(userIcp db.UserIcp) (walletId string, error error) {
	// TODO all balance?
	balance, err := i.getIcpBalance(userIcp.IdentityName)
	if err != nil {
		return "", err
	}
	// create new canister
	createCanisterSprintf := CreateCanister
	createCanisterCmd := fmt.Sprintf(createCanisterSprintf,
		userIcp.PrincipalId, balance, i.network, userIcp.IdentityName)
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
	deployWalletCmd := fmt.Sprintf(deployWalletSprintf, walletId, i.network, userIcp.IdentityName)
	output, err = i.execDfxCommand(deployWalletCmd)
	logger.Infof("userid-> %s walletId-> %s deploy-wallet result is: %s \n", userIcp.IdentityName, walletId, output)
	if err != nil {
		return "", err
	}
	return walletId, nil
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