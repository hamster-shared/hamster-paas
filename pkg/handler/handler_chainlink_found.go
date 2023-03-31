package handler

import (
	"github.com/gin-gonic/gin"
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
		Fail("invalid params", c)
		return
	}

	incrString := c.Query("incr")
	incr, err := strconv.ParseFloat(incrString, 64)
	if err != nil || incr <= 0 {
		Fail("invalid incr", c)
		return
	}

	err = h.chainLinkSubscriptionService.AddFundsForSubscription(subscriptionId, incr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(nil, c)
}
