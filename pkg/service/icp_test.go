package service

import (
	"fmt"
	"hamster-paas/pkg/db"
	"hamster-paas/pkg/utils/logger"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

func getDB() *gorm.DB {
	DSN := fmt.Sprintf("root:%s@tcp(61.172.179.6:30303)/aline?charset=utf8&parseTime=True&loc=Local", os.Getenv("PASSWORD"))
	logger.Debugf("DSN: %s", DSN)
	alineDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DSN,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect aline database: %s", err))
	}
	return alineDb
}

func TestBrief(t *testing.T) {

	logger.InitLogger()
	alineDb := getDB()

	i := IcpService{alineDb, "local"}
	var canisters []db.IcpCanister
	err := i.dbUserCanisters(65406422, &canisters)
	logger.Debugf("ledger info: %v", canisters)
	if err != nil {
		logger.Errorf("ledger info error: %v", err)
	}
}

func TestIdentityName(t *testing.T) {

	logger.InitLogger()
	alineDb := getDB()

	i := IcpService{alineDb, "local"}
	name, err := i.dbIdentityName(65406422)
	logger.Debugf("ledger info: %v", name)
	if err != nil {
		logger.Errorf("ledger info error: %v", err)
	}
}

func TestBalance(t *testing.T) {

	logger.InitLogger()
	// alineDb := getDB()

	i := IcpService{nil, "ic"}
	name, err := i.getCycle("default")
	logger.Debugf("ledger info: %v", name)
	if err != nil {
		logger.Errorf("ledger info error: %v", err)
	}
}
