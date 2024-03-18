package service

import (
	"encoding/json"
	"fmt"
	"hamster-paas/pkg/db"
	"hamster-paas/pkg/utils/logger"
	"io"
	"os"
	"regexp"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestCanisterStatus(t *testing.T) {
	logger.InitLogger()

	i := IcpService{nil, "ic"}
	res, err := i.getCanisterStatus("cold", "we7gr-5iaaa-aaaak-qcxqq-cai")
	logger.Debugf("canister status: %v", res)
	if err != nil {
		logger.Errorf("canister status error: %v", err)
	}
}

func TestBalance(t *testing.T) {
	logger.InitLogger()
	// alineDb := getDB()

	i := IcpService{nil, "ic"}
	name, err := i.getCycle("cold")
	logger.Debugf("ledger info: %v", name)
	if err != nil {
		logger.Errorf("ledger info error: %v", err)
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

// func TestCreate(t *testing.T) {

// 	logger.InitLogger()
// 	alineDb := getDB()

// 	i := IcpService{alineDb, "ic"}
// 	err := i.dbCreateCanister(65406422, "65406422", "nglpe-baaaa-aaaal-qcx3a-cai", "hello")
// 	if err != nil {
// 		logger.Errorf("dbCreateCanister error: %v", err)
// 	}
// }

func TestMatch(t *testing.T) {
	logger.InitLogger()

	text := "hello canister was already created on network ic and has canister id: srbgg-giaaa-aaaao-a27na-cai"
	re := regexp.MustCompile(`has canister id: (.+)`)
	matches := re.FindStringSubmatch(text)
	// logger.Debugf("matches: %v", matches)
	if len(matches) > 1 {
		canisterId := matches[1]
		logger.Debugf("canister id: %s", canisterId)
	} else {
		logger.Errorf("canister id not found")
	}
}

func TestFile(t *testing.T) {
	logger.InitLogger()

	canisterName := "hello"
	// if _, err := os.Stat("dfx.json"); os.IsNotExist(err) {
	// 不存在，则新建并写入数据 {}
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
	// }
}

func TestRead(t *testing.T) {
	logger.InitLogger()

	canisterName := "hello"
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
	logger.Debugf("result: %v", result)
	canisters := result[canisterName].(map[string]interface{})
	logger.Debugf("canisters: %v", canisters["ic"])
}
