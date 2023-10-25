package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "Aline123456", "61.172.179.6", "30303", "paas")
	slowLogger := logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),

		logger.Config{
			//设定慢查询时间阈值为1ms
			SlowThreshold: 15 * time.Millisecond,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel: logger.Warn,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: slowLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_cl_",
		},
	})

	if err != nil {
		panic(err)
	}

	user := aline.User{
		Id: 3366709,
	}

	var zanUser models.ZanUser
	err = db.Model(models.ZanUser{}).Where("user_id =", user.Id).First(&zanUser).Error
	if err != nil {
		fmt.Println(false)
	}
	fmt.Println(zanUser.AccessToken != "")
}

func TestCost(t *testing.T) {
	client, err := ethclient.Dial("http://webtcapi.unchartedw3s.com/node/v1/eth/goerli/f49675d8c5574292ba40adb2a475e770")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		number, err := client.BlockNumber(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(number)
		time.Sleep(time.Second * 10)
	}

}
