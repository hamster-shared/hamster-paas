package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
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
