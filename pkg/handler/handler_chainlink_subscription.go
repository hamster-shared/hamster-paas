package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
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
		UserId:              uint64(user.Id),
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

	primaryId, err := h.chainLinkSubscriptionService.CreateSubscription(s, h.chainlinkPoolService, network.NetworkType())
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
	userAny, exists := gin.Get("user")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	data, err := h.chainLinkSubscriptionService.SubscriptionList(chain, network, page, size, int64(user.Id))
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
	userAny, ok := gin.Get("user")
	if !ok {
		Fail("do not have token", gin)
		return
	}
	user := userAny.(aline.User)

	list, err := h.chainLinkSubscriptionService.GetValidSubscription(int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("get valid subscription failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(list, gin)
}

func (h *HandlerServer) changeSubscriptionStatus(gin *gin.Context) {
	userAny, exists := gin.Get("user")
	if !exists {
		logger.Error(fmt.Sprintf("user not found"))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
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
		Fail("param invalid", gin)
		return
	}
	network, err := models.ParseNetworkType(jsonParam.Network)
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status, network not format: %s", err.Error()))
		Fail("param invalid", gin)
		return
	}
	jsonParam.Chain = chain.String()
	jsonParam.Network = network.StringWithSpace()
	fmt.Println(jsonParam.Chain)
	fmt.Println(jsonParam.Network)
	err = h.chainLinkSubscriptionService.ChangeSubscriptionStatus(jsonParam, uint64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("change subscription status faild: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}
