package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models/vo/node"
	"hamster-paas/pkg/utils/logger"
	"os"
	"strconv"
)

func (h *HandlerServer) nodeList(gin *gin.Context) {
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Errorf("page ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		logger.Errorf("size ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	userId, ok := gin.Get("userId")
	if !ok {
		logger.Errorf("context do not have user")
		Fail("do not have token", gin)
		return
	}
	data, err := h.nodeService.NodeList(userId.(uint), page, size)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) nodeStatisticsInfo(gin *gin.Context) {
	userId, ok := gin.Get("userId")
	if !ok {
		logger.Errorf("context do not have user")
		Fail("do not have token", gin)
		return
	}
	data, err := h.nodeService.NodeStatisticsInfo(userId.(uint))
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) nodeDetail(gin *gin.Context) {
	nodeIdStr := gin.Param("id")
	nodeId, err := strconv.Atoi(nodeIdStr)
	if err != nil {
		logger.Errorf("node id ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	data, err := h.nodeService.NodeDetail(nodeId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) updateNode(gin *gin.Context) {
	_, exists := gin.Get("userId")
	if !exists {
		logger.Errorf("context do not have user")
		Fail("do not have token", gin)
		return
	}

	saveNodeParam := node.UpdateNodeParam{}
	err := gin.BindJSON(&saveNodeParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	if saveNodeParam.VerifyIdentity == os.Getenv("VERIFY_IDENTITY") {
		Fail("You cannot update data", gin)
		return
	}

	err = h.nodeService.UpdateNode(saveNodeParam.ID, saveNodeParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(saveNodeParam, gin)
}

func (h *HandlerServer) launchOrder(gin *gin.Context) {
	userId, ok := gin.Get("userId")
	if !ok {
		logger.Errorf("context do not have user")
		Fail("do not have token", gin)
		return
	}
	launchData := node.LaunchOrderParam{}
	err := gin.BindJSON(&launchData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.orderService.LaunchOrder(userId.(uint), launchData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) orderList(gin *gin.Context) {
	start := gin.GetHeader("X-Start")
	end := gin.GetHeader("X-End")
	query := gin.Query("query")
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Errorf("page ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		logger.Errorf("size ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Errorf("context do not have user")
		Fail("do not have token", gin)
		return
	}
	data, err := h.orderService.OrderList(start, end, query, userId.(uint), page, size)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) payOrderDetail(gin *gin.Context) {
	orderIdStr := gin.Param("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		logger.Errorf("order id ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	data, err := h.orderService.PayOrderDetail(orderId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) cancelOrder(gin *gin.Context) {
	orderIdStr := gin.Param("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		logger.Errorf("order id ato int failed: %s", err)
		Fail(err.Error(), gin)
		return
	}
	err = h.orderService.CancelOrder(orderId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) queryResourceStandard(gin *gin.Context) {
	protocol := gin.Param("protocol")
	data, err := h.resourceStandardService.QueryResourceStandard(protocol, "US East")
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
