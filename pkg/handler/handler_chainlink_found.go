package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"
)

func (h *HandlerServer) addFound(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)
	subscriptionIdString := c.Param("id")
	subscriptionId, err := strconv.Atoi(subscriptionIdString)
	if err != nil {
		logger.Errorf(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail("invalid params", c)
		return
	}
	consumerAddress := c.PostForm("consumerAddress")
	incrString := c.PostForm("incr")
	fmt.Println("incr : ", incrString)
	incr, err := strconv.ParseFloat(incrString, 64)
	if err != nil || incr <= 0 {
		logger.Errorf("addFound failed: incr not valid")
		Fail("invalid incr", c)
		return
	}
	transactionTx := c.PostForm("transactionTx")
	err = h.chainLinkDepositService.AddDeposit(int64(subscriptionId), consumerAddress, incr, transactionTx, int64(user.Id), h.chainLinkSubscriptionService)
	if err != nil {
		logger.Error(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(nil, c)
}
