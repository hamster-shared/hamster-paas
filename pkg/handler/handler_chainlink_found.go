package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models/vo"
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
	foundParam := vo.ChainLinkFoundParam{}
	err = c.BindJSON(&foundParam)
	if err != nil {
		logger.Errorf(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail("invalid incr", c)
		return
	}
	Incr, err := strconv.ParseFloat(foundParam.Incr, 64)
	if err != nil || Incr < 0 {
		logger.Errorf("addFound failed: incr not valid")
		Fail("invalid incr", c)
		return
	}
	err = h.chainLinkDepositService.AddDeposit(int64(subscriptionId), foundParam.ConsumerAddress, Incr, foundParam.TransactionTx, int64(user.Id), h.chainLinkSubscriptionService, h.chainlinkPoolService)
	if err != nil {
		logger.Error(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(nil, c)
}
