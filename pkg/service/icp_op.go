package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"hamster-paas/pkg/db"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"io"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	GetVersion  = "dfx -V"
	NewIdentity = "dfx identity new %s --storage-mode plaintext"
	UseIdentity = "dfx identity use %s"

	AccountId     = "dfx ledger account-id --identity %s"
	LedgerBalance = "dfx ledger balance --network %s --identity %s" // icp
	GetPrincipal  = "dfx identity get-principal --identity %s"
	GetWallet     = "dfx identity get-wallet --network %s --identity %s"
	DeployWallet  = "dfx identity deploy-wallet %s --network %s --identity %s"
	WalletBalance = "dfx wallet balance --network %s --identity %s" // cycle

	WalletTopUp    = "dfx ledger top-up %s --amount %s --network %s --identity %s"
	CreateCanister = "dfx ledger create-canister %s --amount %s --network %s --identity %s"
	DepositCycles  = "dfx canister deposit-cycles %s %s --network %s --identity %s"
	CanisterStatus = "dfx canister status %s --network %s --identity %s"
	CanisterCreate = "dfx canister create %s --with-cycles 200000000000 --network %s --identity %s"
	CanisterDelete = "dfx canister delete %s --network %s --identity %s"
	CanisterStop   = "dfx canister stop %s --network %s --identity %s"
	CanisterStart  = "dfx canister start %s --network %s --identity %s"

	AddController = "dfx canister update-settings %s --add-controller %s --network %s --identity %s"
	DelController = "dfx canister update-settings %s --remove-controller %s --network %s --identity %s"
	UninstallCode = "dfx canister uninstall-code %s --network %s --identity %s"
	InstallCode   = "dfx canister install %s --wasm %s --mode %s --network %s --identity %s -y"

	TransferICP = "dfx ledger transfer %s --icp %s --memo %s --network %s --identity %s"
)

// ################ db operations ################
func (i *IcpService) DBIdentityName(userId uint) (identityName string, err error) {
	var userIcp db.UserIcp
	if err = i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error; err != nil {
		return "", err
	}
	return userIcp.IdentityName, nil
}

func (i *IcpService) dbUserIdentity(userId uint, userIcp *db.UserIcp) error {
	return i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
}

func (i *IcpService) dbUserProjects(userId uint, projects *[]db.Project) error {
	return i.db.Model(db.Project{}).Where("user_id = ?", userId).Find(&projects).Error
}

func (i *IcpService) dbUserCanisters(userId uint, canisters *[]db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("fk_user_id = ?", userId).Order("update_time DESC").Find(&canisters).Error
}

func (i *IcpService) DBAllCanisters(canisters *[]db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Find(&canisters).Error
}

// func (i *IcpService) dbProjCanisters(projId string, canisters *[]db.IcpCanister) error {
// 	return i.db.Model(db.IcpCanister{}).Where("project_id = ?", projId).Find(&canisters).Error
// }

func (i *IcpService) dbProjectInfo(projId string, project *db.Project) error {
	return i.db.Model(db.Project{}).Where("id = ?", projId).First(&project).Error
}

func (i *IcpService) dbCanisterInfo(canisterId string, canister *db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("canister_id = ?", canisterId).First(&canister).Error
}

// create
func (i *IcpService) dbCreateCanister(userId uint, identity string, canisterId string, canisterName string) error {
	out, err := i.getCanisterStatus(identity, canisterId)
	if err != nil {
		return err
	}
	fmt.Printf("out: %v", out)
	canister := db.IcpCanister{
		FkUserId:     userId,
		ProjectId:    "",
		CanisterId:   canisterId,
		CanisterName: canisterName,
		Status:       db.Running,
		Cycles: sql.NullString{
			String: out.Balance,
			Valid:  true,
		},
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	return i.db.Create(&canister).Error
}

func (i *IcpService) dbUpdateCanister(identity string, canisterId string) error {
	out, err := i.getCanisterStatus(identity, canisterId)
	if err != nil {
		return err
	}

	var icpCanister db.IcpCanister
	err = i.db.Model(db.IcpCanister{}).Where("canister_id = ?", canisterId).First(&icpCanister).Error
	if err != nil {
		return err
	}
	icpCanister.Status = db.DBStatus(out.Status)
	icpCanister.Cycles = sql.NullString{
		String: out.Balance,
		Valid:  true,
	}
	icpCanister.UpdateTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	return i.db.Model(db.IcpCanister{}).Where("canister_id = ?", canisterId).Updates(&icpCanister).Error
}

func (i *IcpService) DBSetComsuption(identity string, canisterId string) error {
	out, err := i.getCanisterStatus(identity, canisterId)
	if err != nil {
		return err
	}
	consume := db.IcpConsumption{
		CanisterId: canisterId,
		ModuleHash: out.ModuleHash,
		Cycles: sql.NullString{
			String: out.Balance,
			Valid:  true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	return i.db.Model(db.IcpConsumption{}).Where("canister_id = ?", canisterId).Create(&consume).Error
}

func (i *IcpService) dbGetComsuption(canisterId string, consumptions *[]db.IcpConsumption, page, size int) (int, error) {
	var count int64
	err := i.db.Model(db.IcpConsumption{}).Where("canister_id = ?", canisterId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	tx := i.db.Model(db.IcpConsumption{}).Where("canister_id = ?", canisterId)
	err = tx.Order("update_time DESC").Offset((page - 1) * size).Limit(size).Find(&consumptions).Error
	return int(count), err
}

// delete
func (i *IcpService) dbDeleteCanister(userId uint, canisterId string) error {
	return i.db.Delete(db.IcpCanister{}).Where("canister_id = ?", canisterId).Error
}

//################ dfx operations ################

func (i *IcpService) newIndentity(identity string) (err error) {
	newIdentityCmd := fmt.Sprintf(NewIdentity, identity)
	_, err = i.execDfxCommand(newIdentityCmd)
	return err
}

// func (i *IcpService) useIndentity(identity string) (err error) {
// 	useIdentityCmd := fmt.Sprintf(UseIdentity, identity)
// 	_, err = i.execDfxCommand(useIdentityCmd)
// 	return err
// }

// return accountId, principal
func (i *IcpService) getLedgerInfo(identity string) (string, string, error) {
	accountIdCmd := fmt.Sprintf(AccountId, identity)
	accountId, err := i.execDfxCommand(accountIdCmd)
	if err != nil {
		return "", "", err
	}
	accountId = strings.TrimSpace(accountId)
	principalCmd := fmt.Sprintf(GetPrincipal, identity)
	principal, err := i.execDfxCommand(principalCmd)
	if err != nil {
		return "", "", err
	}
	principal = strings.TrimSpace(principal)
	return accountId, principal, nil
}

func (i *IcpService) getWalletId(identity string) (string, error) {
	getWalletCmd := fmt.Sprintf(GetWallet, i.network, identity)
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

// RAW dfx ledger balance --network ic
func (i *IcpService) getIcp(identity string) (string, error) {
	balanceCmd := fmt.Sprintf(LedgerBalance, i.network, identity)
	balance, err := i.execDfxCommand(balanceCmd)
	if err != nil {
		return "", err
	}
	balanceSplit := strings.Split(balance, " ")
	if len(balanceSplit) > 0 {
		return balanceSplit[0], nil
	} else {
		return "", errors.New("failure to obtain icp balances")
	}
}

// RAW dfx wallet balance --network ic
func (i *IcpService) getCycle(identity string) (string, error) {
	walletBalanceCmd := fmt.Sprintf(WalletBalance, i.network, identity)
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

func (i *IcpService) getIcpBalance(identity string) (string, error) {
	balanceCmd := fmt.Sprintf(LedgerBalance, i.network, identity)
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

// top up all ICPs
func (i *IcpService) walletTopUp(identity string, walletId string) error {
	// TODO all balance topup?
	balance, err := i.getIcpBalance(identity)
	if err != nil {
		return err
	}
	walletTopUpCmd := fmt.Sprintf(WalletTopUp, walletId, balance, i.network, identity)
	output, err := i.execDfxCommand(walletTopUpCmd)
	if err != nil {
		return err
	}
	logger.Infof("identityName-> %s walletId-> %s top-up result is: %s \n", identity, walletId, output)
	return nil
}

// deposit cycles
func (i *IcpService) depositCanister(identity string, cycles string, canisterId string) error {
	depositCyclesCmd := fmt.Sprintf(DepositCycles, cycles, canisterId, i.network, identity)
	output, err := i.execDfxCommand(depositCyclesCmd)
	if err != nil {
		return err
	}
	logger.Infof("userid-> %s canisterId-> %s deposit-cycles result is: %s \n", identity, canisterId, output)
	return nil
}

// status all
func (i *IcpService) getCanisterStatus(identity string, canisterId string) (*vo.CanisterStatus, error) {
	var res vo.CanisterStatus
	statusCmd := fmt.Sprintf(CanisterStatus, canisterId, i.network, identity)
	logger.Infof("exec cmd is %s", statusCmd)
	cmd := exec.Command("bash", "-c", statusCmd)
	out, err := cmd.CombinedOutput()
	result := string(out)
	if err != nil {
		return nil, err
	}
	// logger.Infof("status is:%s", result)
	// status
	re := regexp.MustCompile(`Status: (.+)`)
	matches := re.FindStringSubmatch(result)
	if len(matches) > 1 {
		res.Status = matches[1]
	} else {
		logger.Errorf("canister status not found")
	}
	// controllers
	re = regexp.MustCompile(`Controllers: (.+)`)
	matches = re.FindStringSubmatch(result)
	if len(matches) > 1 {
		res.Controllers = strings.Split(matches[1], " ")
	} else {
		logger.Errorf("controllers not found!")
	}
	// cycles
	re = regexp.MustCompile(`Balance: ([0-9_]+) Cycles`)
	matches = re.FindStringSubmatch(string(result))
	if len(matches) > 1 {
		value := matches[1]
		value = strings.ReplaceAll(value, "_", "")
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Errorf("balance parse int failed:%s", err)
			return nil, err
		}
		data := float64(number) / math.Pow(10, 12)
		balance := fmt.Sprintf("%.2f", data)
		res.Balance = balance
	} else {
		logger.Errorf("balance not found!")
	}
	// memory size
	re = regexp.MustCompile(`Memory Size: Nat\((.+)\)`)
	matches = re.FindStringSubmatch(result)
	if len(matches) > 1 {
		value := matches[1]
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Errorf("memorysize parse int failed:%s", err)
			return nil, err
		}
		data := float64(number*16) / math.Pow(2, 20)
		size := fmt.Sprintf("%.2f", data)
		res.MemorySize = size
	} else {
		logger.Errorf("memory size not found!")
	}
	// module hash
	re = regexp.MustCompile(`Module hash: (.+)`)
	matches = re.FindStringSubmatch(result)
	if len(matches) > 1 {
		res.ModuleHash = matches[1]
	} else {
		logger.Errorf("module hash not found!")
	}
	// logger.Debugf("canister status result is %v", res)

	return &res, err
}

// 1 running 0 or 2 stoped
func (i *IcpService) changeCanisterStatus(identity string, canisterId string, statusType vo.StatusType) error {
	if _, err := os.Stat("dfx.json"); os.IsNotExist(err) {
		// 不存在，则新建并写入数据 {}
		data := map[string]interface{}{}
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = os.WriteFile("dfx.json", dataJSON, 0644)
		if err != nil {
			return err
		}
	}
	var changStatusCmd string
	if statusType == vo.Running {
		changStatusCmd = fmt.Sprintf(CanisterStart, canisterId, i.network, identity)
		output, err := i.execDfxCommand(changStatusCmd)
		if err != nil {
			return err
		}
		logger.Infof("userid-> %s canisterId-> %s start result is: %s \n", identity, canisterId, output)

	} else {
		changStatusCmd = fmt.Sprintf(CanisterStop, canisterId, i.network, identity)
		output, err := i.execDfxCommand(changStatusCmd)
		if err != nil {
			return err
		}
		logger.Infof("userid-> %s canisterId-> %s stop result is: %s \n", identity, canisterId, output)
	}

	return nil
}

func (i *IcpService) createCanister(identity string, canisterName string) (canisterId string, err error) {
	// 写入数据
	data := map[string]interface{}{}
	data["canisters"] = map[string]interface{}{canisterName: map[string]interface{}{}}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		logger.Errorf("marshal dfx.json error: %v", err)
	}
	err = os.WriteFile("dfx.json", dataJSON, 0644)
	if err != nil {
		logger.Errorf("write dfx.json error: %v", err)
	}
	//创建 canister
	createCanisterCmd := fmt.Sprintf(CanisterCreate, canisterName, i.network, identity)
	out, err := i.execDfxCommand(createCanisterCmd)
	logger.Debugf("create canister: \ncmd %s \nout %s", createCanisterCmd, out)

	if err != nil {
		return "", err
	}
	// re := regexp.MustCompile(`Canister created with id: "(.*?)"`)
	re := regexp.MustCompile(`has canister id: (.+)`)
	matches := re.FindStringSubmatch(out)
	if len(matches) > 1 {
		canisterId = matches[1]
	} else {
		jsonFile, err := os.Open("canister_ids.json")
		if err != nil {
			logger.Errorf("open canister_ids.json error: %v", err)
		}
		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			logger.Errorf("read canister_ids.json error: %v", err)
		}
		var result map[string]interface{}
		json.Unmarshal(byteValue, &result)
		// logger.Debugf("result: %v", result)
		canisters := result[canisterName].(map[string]interface{})
		canisterId = canisters["ic"].(string)
		// logger.Debugf("canisters: %v", canisterId)
		if canisterId == "" {
			return "", errors.New("canister id not found")
		} else {
			return canisterId, nil
		}
	}
	logger.Infof("identity-> %s canisterName-> %s create-canister result is: %s \n", identity, canisterName, canisterId)
	return canisterId, nil
}

func (i *IcpService) deleteCanister(identity string, canisterId string) error {
	deleteCanisterCmd := fmt.Sprintf(CanisterDelete, canisterId, i.network, identity)
	out, err := i.execDfxCommand(deleteCanisterCmd)
	logger.Debugf("delete canister: \ncmd %s \nout %s", deleteCanisterCmd, out)

	if err != nil {
		return err
	}
	logger.Infof("canisterId-> %s delete-canister result is: %s \n", canisterId, out)
	return nil
}

func (i *IcpService) addController(identity string, canisterId string, controller string) error {
	addControllerCmd := fmt.Sprintf(AddController, canisterId, controller, i.network, identity)
	output, err := i.execDfxCommand(addControllerCmd)
	logger.Debugf("add controller: \ncmd %s \nout %s", addControllerCmd, output)

	if err != nil {
		return err
	}
	logger.Infof("userid-> %s canisterId-> %s add-controller %s result is: %s \n", identity, canisterId, controller, output)
	return nil
}

func (i *IcpService) delController(identity string, canisterId string, controller string) error {
	delControllerCmd := fmt.Sprintf(DelController, canisterId, controller, i.network, identity)
	output, err := i.execDfxCommand(delControllerCmd)
	logger.Debugf("del controller: \ncmd %s \nout %s", delControllerCmd, output)

	if err != nil {
		return err
	}
	logger.Infof("userid-> %s canisterId-> %s del-controller %s result is: %s \n", identity, canisterId, controller, output)
	return nil
}

func (i *IcpService) installWasm(identity string, canisterId string, wasmPath string, mode vo.InstallMode) error {
	if _, err := os.Stat("dfx.json"); os.IsNotExist(err) {
		// 不存在，则新建并写入数据 {}
		data := map[string]interface{}{}
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = os.WriteFile("dfx.json", dataJSON, 0644)
		if err != nil {
			return err
		}
	}
	installWasmCmd := fmt.Sprintf(InstallCode, canisterId, wasmPath, mode.String(), i.network, identity)
	output, err := i.execDfxCommand(installWasmCmd)
	logger.Debugf("install wasm: \ncmd %s \nout %s", installWasmCmd, output)
	if err != nil {
		return err
	}
	logger.Infof("canisterId-> %s install-wasm %s result is: %s \n", canisterId, wasmPath, output)
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
