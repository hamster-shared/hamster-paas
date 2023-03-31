package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	subscriptionIdString := c.Query("subscription_id")
	subscriptionId, err := strconv.Atoi(subscriptionIdString)
	if err != nil {
		Fail("invalid params", c)
		return
	}

	consumer := models.Consumer{
		SubscriptionId: int64(subscriptionId),
		Created:        time.Now(),
		UserId:         uint64(user.Id),
	}

	// 确保订阅存在
	subscription, err := h.chainLinkSubscriptionService.GetSubscriptionById(subscriptionId)
	if err != nil {
		Fail(err.Error(), c)
		return
	}

	// 创建合约
	consumerNums := h.chainLinkConsumerService.CreateConsumer(consumer, subscriptionId)
	// 往Subscription增加合约数量
	if err := h.chainLinkSubscriptionService.AddConsumer(subscription.SubscriptionId, consumerNums); err != nil {
		Fail(err.Error(), c)
		return
	}

	Success(nil, c)
}

// TODO: 暂时直接返回假数据
func (h *HandlerServer) getConsumerList(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		Fail("invalid params", c)
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		Fail("invalid params", c)
		return
	}
	var pagination models.Pagination
	pagination.Page = pageInt
	pagination.Size = sizeInt
	// 查询hamster的可用合约列表
	var consumerList []vo.ChainLinkConsumers
	if pagination.Page > 1 {
		SuccessWithPagination(consumerList, pagination, c)
		return
	}
	consumerList = append(consumerList, vo.ChainLinkConsumers{
		Address:    "0x123456789",
		Network:    "Test",
		DeployTime: time.Now(),
	})
	consumerList = append(consumerList, vo.ChainLinkConsumers{
		Address:    "0x123456789",
		Network:    "Test",
		DeployTime: time.Now(),
	})
	consumerList = append(consumerList, vo.ChainLinkConsumers{
		Address:    "0x123456789",
		Network:    "Test",
		DeployTime: time.Now(),
	})
	SuccessWithPagination(consumerList, pagination, c)
	return
}

func (h *HandlerServer) consumerList(gin *gin.Context) {
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
	data, err := h.chainLinkConsumerService.ConsumerList(page, size, int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("query consumer list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) deleteConsumer(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("chainlink consumer id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	err = h.chainLinkConsumerService.DeleteConsumer(int64(id))
	if err != nil {
		logger.Error(fmt.Sprintf("delete consumer failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}
