package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"strconv"
)

func (h *HandlerServer) addFound(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
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
	primaryId, err := h.chainLinkDepositService.AddDeposit(int64(subscriptionId), Incr, foundParam.TransactionTx, userId.(uint), h.chainLinkSubscriptionService, h.chainlinkPoolService, foundParam.Address)
	if err != nil {
		logger.Error(fmt.Sprintf("addFound failed: %s", err.Error()))
		Fail(err.Error(), c)
		return
	}
	Success(primaryId, c)
}

func (h *HandlerServer) changeFoundStatus(gin *gin.Context) {
	userId, exists := gin.Get("userId")
	if !exists {
		Fail("do not have token", gin)
		return
	}
	var paramJson vo.ChainLinkFoundUpdateParam
	if err := gin.BindJSON(&paramJson); err != nil {
		logger.Error(fmt.Sprintf("change Found Status failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	status, err := consts.ParseStatus(paramJson.NewStatus)
	if err != nil {
		logger.Error(fmt.Sprintf("change Found Status failed: status invalid %s", err.Error()))
		Fail(fmt.Sprintf("change Found Status failed: status invalid %s", err.Error()), gin)
		return
	}
	paramJson.NewStatus = status
	err = h.chainLinkDepositService.UpdateDepositStatus(userId.(uint), paramJson)
	if err != nil {
		logger.Error(fmt.Sprintf("change Found Status failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(nil, gin)
}
