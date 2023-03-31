package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/utils/logger"
	"strconv"
)

func (h *HandlerServer) addFound(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	subscriptionIdString := c.Query("subscription_id")
	subscriptionId, err := strconv.Atoi(subscriptionIdString)
	if err != nil {
		logger.Errorf(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail("invalid params", c)
		return
	}
	incrString := c.Query("incr")
	incr, err := strconv.ParseFloat(incrString, 64)
	if err != nil || incr <= 0 {
		logger.Errorf("addFound failed: incr not valid")
		Fail("invalid incr", c)
		return
	}
	err = h.chainLinkSubscriptionService.AddFundsForSubscription(subscriptionId, incr)
	if err != nil {
		logger.Error(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(nil, c)
}
