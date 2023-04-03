package service

import (
	"hamster-paas/pkg/utils/logger"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	logger.InitLogger()

	pool := NewPoolService()

	pool.Submit(func() {
		logger.Info("start task")
		time.Sleep(time.Second * 10)
		logger.Info("stop task")
	})
	logger.Info("start wait")
	pool.Close()
}
