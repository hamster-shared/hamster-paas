package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"
	"time"
)

func (h *HandlerServer) createConsumer(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)
	consumerCreateParam := vo.ChainLinkConsumerCreateParam{}
	err := c.BindJSON(&consumerCreateParam)
	if err != nil {
		logger.Error(fmt.Sprintf("add consumer params vaild %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	consumer := models.Consumer{
		SubscriptionId:  consumerCreateParam.SubscriptionId,
		Created:         time.Now(),
		ConsumerAddress: consumerCreateParam.ConsumerAddress,
		UserId:          uint64(user.Id),
		TransactionTx:   consumerCreateParam.TransactionTx,
		Status:          consts.PENDING,
	}
	// 创建合约
	primaryId, err := h.chainLinkConsumerService.CreateConsumer(consumer, h.chainLinkSubscriptionService, h.chainlinkPoolService)
	if err != nil {
		logger.Error(fmt.Sprintf("add consumer in subscriptionL: %d failed: %s", consumerCreateParam.SubscriptionId, err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(primaryId, c)
}

func (h *HandlerServer) getHamsterConsumerList(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		Fail("invalid params: page", c)
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		Fail("invalid params: size", c)
		return
	}
	var pagination models.Pagination
	pagination.Page = pageInt
	pagination.Size = sizeInt
	projectIdString := c.Param("id")
	chainParam := c.Query("chain")
	networkParam := c.Query("network")
	if chainParam == "" && networkParam == "" {
		Fail("invalid params: chain an network", c)
		return
	}
	network := models.GetAlineNetwork(chainParam, networkParam)
	// 通过project id 拿到对应的 consumer 信息
	projectService, err := application.GetBean[*aline.ProjectService]("projectService")
	if err != nil {
		Fail("get project service error", c)
		return
	}
	data, err := projectService.GetValidContract(pageInt, sizeInt, projectIdString, network)
	if err != nil {
		logger.Error(fmt.Sprintf("get Hamster Consumer List failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

func (h *HandlerServer) consumerList(gin *gin.Context) {
	idStr := gin.Param("id")
	subscriptionId, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("subscription id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
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
		logger.Error(fmt.Sprintf("user not found: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	data, err := h.chainLinkConsumerService.ConsumerList(subscriptionId, page, size, int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("query consumer list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) consumerAddressList(gin *gin.Context) {
	idStr := gin.Param("id")
	subscriptionId, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("subscription id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	userAny, exists := gin.Get("user")
	if !exists {
		logger.Error(fmt.Sprintf("user not found: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	data, err := h.chainLinkConsumerService.ConsumerAddressList(int64(subscriptionId), int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("query consumer list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) deleteConsumer(gin *gin.Context) {
	idStr := gin.Param("id")
	subscriptionId, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("chainlink subscription id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	consumerIdStr := gin.Param("consumerId")
	consumerId, err := strconv.Atoi(consumerIdStr)
	if err != nil {
		logger.Error(fmt.Sprintf("chainlink consumer id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	err = h.chainLinkConsumerService.DeleteConsumer(int64(subscriptionId), int64(consumerId))
	if err != nil {
		logger.Error(fmt.Sprintf("delete consumer failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) changeConsumerStatus(gin *gin.Context) {
	userAny, exists := gin.Get("user")
	if !exists {
		logger.Error(fmt.Sprintf("user not found"))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	var jsonParam vo.ChainLinkConsumerUpdateParam
	err := gin.BindJSON(&jsonParam)
	if err != nil {
		logger.Error(fmt.Sprintf("change consumer status: param invalid: %s", err.Error()))
		Fail("param invalid", gin)
		return
	}
	status, err := consts.ParseStatus(jsonParam.NewStatus)
	if err != nil {
		logger.Error(fmt.Sprintf("change consumer status: param invalid: %s", err.Error()))
		Fail(fmt.Sprintf("change consumer status: status invalid: %s", err.Error()), gin)
		return
	}
	jsonParam.NewStatus = status
	err = h.chainLinkConsumerService.ChangeConsumerStatus(jsonParam, uint64(user.Id), h.chainLinkSubscriptionService)
	if err != nil {
		logger.Error(fmt.Sprintf("change consumer status faild: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}
