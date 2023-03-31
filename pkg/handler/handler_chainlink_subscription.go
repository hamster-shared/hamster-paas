package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"log"
	"strconv"
	"time"
)

func (h *HandlerServer) getSubscriptionOverview(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)

	network := c.Query("network")
	if network == "" {
		Fail("network not valid", c)
		return
	}

	ov, err := h.chainLinkSubscriptionService.GetSubscriptionOverview(user.Id, network)
	if err != nil {
		log.Println(err)
		Fail(err.Error(), c)
		return
	}

	Success(ov, c)
}

func (h *HandlerServer) getSINA(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)

	sinas := h.chainLinkSubscriptionService.GetSINAByUserId(user.Id)
	Success(sinas, c)
}

func (h *HandlerServer) createSubscription(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)

	// get chain
	chain := c.Query("chain")
	if chain == "" {
		Fail("chain not valid", c)
		return
	}
	// get network
	network := c.Query("network")
	if network == "" {
		Fail("network not valid", c)
		return
	}
	// get name
	name := c.Query("name")
	if name == "" {
		Fail("name not valid", c)
		return
	}
	// get id
	Id := c.Query("subscriptionId")
	subscriptionId, err := strconv.Atoi(Id)
	if err != nil {
		Fail("invalid params", c)
		return
	}

	s := models.Subscription{
		SubscriptionId: uint(subscriptionId),
		Name:           name,
		Chain:          chain,
		Network:        network,
		UserId:         uint64(user.Id),
		Created:        time.Now(),
	}

	if err := h.chainLinkSubscriptionService.CreateSubscription(s); err != nil {
		logger.Error(fmt.Sprintf("Create subscription failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}

	Success(nil, c)
}
