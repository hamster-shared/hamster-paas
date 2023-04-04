package aline

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"hamster-paas/pkg/application"
	"os"
)

type AlineRpc struct {
}

func NewAlineRpc() *AlineRpc {
	return &AlineRpc{}
}

func (a *AlineRpc) Init() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("ALINE_DB_USER"), os.Getenv("ALINE_DB_PASSWORD"), os.Getenv("ALINE_DB_HOST"), os.Getenv("ALINE_DB_PORT"), os.Getenv("ALINE_DB_NAME"))
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
	application.SetBean[*gorm.DB]("alineDb", alineDb)
	userService := NewUserService()
	application.SetBean[*UserService]("userService", userService)
	projectService := NewAlineProjectService(alineDb)
	application.SetBean[*ProjectService]("projectService", projectService)
}
