package service

import (
	"hamster-paas/pkg/utils/logger"
	"testing"
)

func TestCanisterStatus(t *testing.T) {

	logger.InitLogger()

	i := IcpService{nil, "local"}
	res, err := i.getCanisterStatus("default", "bd3sg-teaaa-aaaaa-qaaba-cai")
	logger.Debugf("canister status: %v", res)
	if err != nil {
		logger.Errorf("canister status error: %v", err)
	}
}

func TestLedgerInfo(t *testing.T) {

	logger.InitLogger()

	i := IcpService{nil, "local"}
	a, b, err := i.getLedgerInfo("default")
	logger.Debugf("ledger info: %s %s", a, b)
	if err != nil {
		logger.Errorf("ledger info error: %v", err)
	}
}
