package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"
	"time"
)

func (h *HandlerServer) getSubscriptionOverview(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("get user info error", c)
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
		logger.Error(fmt.Sprintf("getSubscriptionOverview failed: %s", err.Error()))
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

// TODO: 先存db，type = Pending，异步查Tx，TX正确，修改type=Success
func (h *HandlerServer) createSubscription(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)
	chain := c.PostForm("chain")
	network := c.PostForm("network")
	name := c.PostForm("name")
	idString := c.PostForm("subscriptionId")
	subscriptionId, err := strconv.Atoi(idString)
	if err != nil {
		logger.Error(fmt.Sprintf("createSubscription failed: %s", err.Error()))
		Fail("invalid params", c)
		return
	}
	admin := c.PostForm("admin")
	transactionTx := c.PostForm("transactionTx")
	s := models.Subscription{
		ChainSubscriptionId: uint(subscriptionId),
		Name:                name,
		Created:             time.Now(),
		Chain:               chain,
		Network:             network,
		Consumers:           0,
		UserId:              uint64(user.Id),
		Admin:               admin,
		TransactionTx:       transactionTx,
		Status:              "Pending",
	}
	if err := h.chainLinkSubscriptionService.CreateSubscription(s); err != nil {
		logger.Error(fmt.Sprintf("Create subscription failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	// TODO: 异步查状态，如果tx正确，修改type = Success
	Success(nil, c)
}

func (h *HandlerServer) subscriptionList(gin *gin.Context) {
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	network := gin.Query("network")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	userAny, exists := gin.Get("user")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	data, err := h.chainLinkSubscriptionService.SubscriptionList(network, page, size, int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) subscriptionDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("chainlink subscription id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	data, err := h.chainLinkSubscriptionService.SubscriptionDetail(int64(id))
	if err != nil {
		logger.Error(fmt.Sprintf("get subscription detail failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
