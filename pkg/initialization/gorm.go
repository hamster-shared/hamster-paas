package initialization

import (
	"fmt"
	"hamster-paas/pkg/application"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_cl_",
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}
	application.SetBean("db", db)
}
