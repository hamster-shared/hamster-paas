package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"strconv"
	"time"
)

func (h *HandlerServer) getSubscriptionOverview(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("get user info error", c)
		return
	}
	network := c.Query("network")
	if network == "" {
		Fail("network not valid", c)
		return
	}
	networkType, err := models.ParseNetworkType(network)
	if err != nil {
		Fail(fmt.Sprintf("network not valid: %s", network), c)
		return
	}
	ov, err := h.chainLinkSubscriptionService.GetSubscriptionOverview(userId.(uint), networkType.String())
	if err != nil {
		logger.Error(fmt.Sprintf("getSubscriptionOverview failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(ov, c)
}

func (h *HandlerServer) getSINA(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	sinas := h.chainLinkSubscriptionService.GetSINAByUserId(userId.(uint))
	Success(sinas, c)
}

// TODO: 先存db，type = Pending，异步查Tx，TX正确，修改type=Success
func (h *HandlerServer) createSubscription(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	subscriptionCreateParam := vo.ChainLinkSubscriptionCreateParam{}
	err := c.BindJSON(&subscriptionCreateParam)
	if err != nil {
		logger.Error(fmt.Sprintf("Create subscription failed, bind json error: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	s := models.Subscription{
		ChainSubscriptionId: uint(subscriptionCreateParam.SubscriptionId),
		Name:                subscriptionCreateParam.Name,
		Created:             time.Now(),
		Chain:               subscriptionCreateParam.Chain,
		Network:             subscriptionCreateParam.Network,
		Consumers:           0,
		UserId:              uint64(userId.(uint)),
		Admin:               subscriptionCreateParam.Admin,
		TransactionTx:       subscriptionCreateParam.TransactionTx,
		Status:              consts.PENDING,
	}
	chain, err := models.ParseChainType(s.Chain)
	if err != nil {
		logger.Error(fmt.Sprintf("chain format error: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	network, err := models.ParseNetworkType(s.Network)
	if err != nil {
		logger.Error(fmt.Sprintf("network format error: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	s.Chain = chain.String()
	s.Network = network.StringWithSpace()

	primaryId, err := h.chainLinkSubscriptionService.CreateSubscription(s, h.chainlinkPoolService, network)
	if err != nil {
		logger.Error(fmt.Sprintf("Create subscription failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(primaryId, c)
}

func (h *HandlerServer) subscriptionList(gin *gin.Context) {
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	chain := gin.Query("chain")
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
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	data, err := h.chainLinkSubscriptionService.SubscriptionList(chain, network, page, size, userId.(uint))
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

func (h *HandlerServer) getValidSubscription(gin *gin.Context) {
	userId, exists := gin.Get("userId")
	if !exists {
		Fail("do not have token", gin)
		return
	}

	list, err := h.chainLinkSubscriptionService.GetValidSubscription(userId.(uint))
	if err != nil {
		logger.Error(fmt.Sprintf("get valid subscription failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(list, gin)
}

func (h *HandlerServer) changeSubscriptionStatus(gin *gin.Context) {
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("user not found"))
		Fail("user information does not exist", gin)
		return
	}
	var jsonParam vo.ChainLinkSubscriptionUpdateParam
	err := gin.BindJSON(&jsonParam)
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status: param invalid: %s", err.Error()))
		Fail("param invalid", gin)
		return
	}
	chain, err := models.ParseChainType(jsonParam.Chain)
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status, chain not format: %s", err.Error()))
		Fail(fmt.Sprintf("change subscription status, chain not format: %s", err.Error()), gin)
		return
	}
	network, err := models.ParseNetworkType(jsonParam.Network)
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status, network not format: %s", err.Error()))
		Fail(fmt.Sprintf("change subscription status, network not format: %s", err.Error()), gin)
		return
	}
	status, err := consts.ParseStatus(jsonParam.NewStatus)
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status, status not format: %s", err.Error()))
		Fail(fmt.Sprintf("change subscription status, status not format: %s", err.Error()), gin)
		return
	}
	jsonParam.Chain = chain.String()
	jsonParam.Network = network.StringWithSpace()
	jsonParam.NewStatus = status
	err = h.chainLinkSubscriptionService.ChangeSubscriptionStatus(jsonParam, userId.(uint))
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status faild: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}

// getSubscriptionBalanceAll
func (h *HandlerServer) getSubscriptionBalanceAll(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	balance, err := h.chainLinkSubscriptionService.GetUserSubscriptionBalanceAll(userId.(uint))
	if err != nil {
		logger.Error(fmt.Sprintf("getSubscriptionBalanceAll failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(balance, c)
}

// getSubscriptionBalanceById
func (h *HandlerServer) getSubscriptionBalanceById(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("getSubscriptionBalanceById failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	balance, err := h.chainLinkSubscriptionService.GetUserSubscriptionBalanceById(userId.(uint), uint64(id))
	if err != nil {
		logger.Error(fmt.Sprintf("getSubscriptionBalanceById failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(balance, c)
}
