package service

import (
	"errors"
	"fmt"
	"hamster-paas/pkg/db"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	GetVersion     = "dfx -V"
	NewIdentity    = "dfx identity new %s --storage-mode plaintext"
	UseIdentity    = "dfx identity use %s"
	AccountId      = "dfx ledger account-id --identity %s"
	GetPrincipal   = "dfx identity get-principal --identity %s"
	GetWallet      = "dfx identity get-wallet --network %s --identity %s"
	DeployWallet   = "dfx identity deploy-wallet %s --network %s --identity %s"
	LedgerBalance  = "dfx ledger balance --network %s --identity %s" // icp
	WalletBalance  = "dfx wallet balance --network %s --identity %s" // cycle
	CreateCanister = "dfx ledger create-canister %s --amount %s --network %s --identity %s"
	WalletTopUp    = "dfx ledger top-up %s --amount %s --network %s --identity %s"
	DepositCycles  = "dfx canister deposit-cycles %s %s --network %s --identity %s"
	CanisterStatus = "dfx canister status %s --network %s --identity %s"
	CanisterCreate = "dfx canister create %s --all --network %s --identity %s"
	AddController  = "dfx canister update-settings --add-controller %s %s --network %s --identity %s"
	DelController  = "dfx canister update-settings --remove-controller %s %s --network %s --identity %s"
	UninstallCode  = "dfx canister uninstall-code %s --network %s --identity %s"
	TransferICP    = "dfx ledger transfer %s --icp %s --memo %s --network %s --identity %s"
)

// db operations
func (i *IcpService) dbUserIdentity(userId uint, userIcp *db.UserIcp) error {
	return i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
}

func (i *IcpService) dbUserProjects(userId uint, projects *[]db.Project) error {
	return i.db.Model(db.Project{}).Where("user_id = ?", userId).Find(&projects).Error
}

func (i *IcpService) dbUserCanisters(userId uint, canisters *[]db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("fk_user_id = ?", userId).Find(&canisters).Error
}

func (i *IcpService) dbProjCanisters(projId string, canisters *[]db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("project_id = ?", projId).Find(&canisters).Error
}

func (i *IcpService) dbProjectInfo(projId string, project *db.Project) error {
	return i.db.Model(db.Project{}).Where("project_id = ?", projId).First(&project).Error
}

func (i *IcpService) dbCanisterInfo(canisterId string, canister *db.IcpCanister) error {
	return i.db.Model(db.IcpCanister{}).Where("canister_id = ?", canisterId).First(&canister).Error
}

// dfx operations
func (i *IcpService) newIndentity(identity string) (err error) {
	newIdentityCmd := fmt.Sprintf(NewIdentity, identity)
	_, err = i.execDfxCommand(newIdentityCmd)
	return err
}

func (i *IcpService) useIndentity(identity string) (err error) {
	useIdentityCmd := fmt.Sprintf(UseIdentity, identity)
	_, err = i.execDfxCommand(useIdentityCmd)
	return err
}

// return accountId, principal
func (i *IcpService) getLedgerInfo(identity string) (string, string, error) {
	accountIdCmd := fmt.Sprintf(AccountId, identity)
	accountId, err := i.execDfxCommand(accountIdCmd)
	if err != nil {
		return "", "", err
	}
	principalCmd := fmt.Sprintf(GetPrincipal, identity)
	principal, err := i.execDfxCommand(principalCmd)
	if err != nil {
		return "", "", err
	}

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

func (i *IcpService) icpBalanceWithUnit(identity string) (string, error) {
	ledgerBalanceCmd := fmt.Sprintf(LedgerBalance, i.network, identity)
	balance, err := i.execDfxCommand(ledgerBalanceCmd)
	return strings.TrimSpace(balance), err
}

func (i *IcpService) cycleBalanceWithUnit(identity string) (string, error) {
	walletBalanceCmd := fmt.Sprintf(WalletBalance, i.network, identity)
	balance, err := i.execDfxCommand(walletBalanceCmd)
	return strings.TrimSpace(balance), err
}

// deprecated
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

func (i *IcpService) depositCanister(identity string, cycles string, canisterId string) error {
	depositCyclesCmd := fmt.Sprintf(DepositCycles, cycles, canisterId, i.network, identity)
	output, err := i.execDfxCommand(depositCyclesCmd)
	if err != nil {
		return err
	}
	logger.Infof("userid-> %s canisterId-> %s deposit-cycles result is: %s \n", identity, canisterId, output)
	return nil
}

func (i *IcpService) getCanisterStatus(identity string, canisterId string) (*vo.CanisterStatus, error) {
	var res vo.CanisterStatus
	statusCmd := fmt.Sprintf(CanisterStatus, canisterId, i.network, identity)
	logger.Infof("exec cmd is %s", statusCmd)
	cmd := exec.Command("bash", "-c", statusCmd)
	out, err := cmd.CombinedOutput()
	result := string(out)
	// logger.Debugf("canister status result is %s", result)
	if err != nil {
		return nil, err
	}
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
	return &res, err
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
