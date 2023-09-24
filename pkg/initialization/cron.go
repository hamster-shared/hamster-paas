package initialization

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func TestCron() {
	//cron.WithSeconds()	秒级操作
	//cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger))	函数没执行完就跳过本次函数
	//cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))	打印任务日志
	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	i := 1
	EntryID, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println(time.Now(), "每5s一次----------------", i)
		time.Sleep(time.Second * 6)
		i++
	})
	fmt.Println(time.Now(), EntryID, err)

	c.Start()
}
