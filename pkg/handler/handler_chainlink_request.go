package handler

import (
	"fmt"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) requestList(gin *gin.Context) {
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
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	data, err := h.chainLinkRequestService.RequestList(page, size, userId.(uint))
	if err != nil {
		logger.Error(fmt.Sprintf("get request list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) saveChainLinkRequest(gin *gin.Context) {
	createData := vo.ChainLinkRequestParam{}
	err := gin.BindJSON(&createData)
	if err != nil {
		logger.Error(fmt.Sprintf("create chainlink request failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	data := vo.ChainLinkRequest{
		Name:        createData.Name,
		Script:      createData.Script,
		UserId:      uint64(userId.(uint)),
		ParamsCount: createData.ParamsCount,
	}
	err = h.chainLinkRequestService.SaveChainLinkRequest(data)
	if err != nil {
		logger.Error(fmt.Sprintf("save chainlink request failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) updateChainLinkRequest(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("chainlink request id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	var updateData vo.ChainLinkRequestParam
	err = gin.BindJSON(&updateData)
	if err != nil {
		logger.Error(fmt.Sprintf("update param question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	data := vo.ChainLinkRequest{
		Name:   updateData.Name,
		Script: updateData.Script,
		UserId: uint64(userId.(uint)),
	}
	err = h.chainLinkRequestService.UpdateChainLinkRequest(int64(id), data)
	if err != nil {
		logger.Error(fmt.Sprintf("update chainlink request failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) chainLinkRequestTemplateList(gin *gin.Context) {
	data, err := h.chainLinkRequestService.ChainLinkRequestTemplateList()
	if err != nil {
		logger.Error(fmt.Sprintf("get chainlink request template failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) getRequestTemplateScript(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("template id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	data, err := h.chainLinkRequestService.GetRequestTemplateScript(int64(id))
	if err != nil {
		logger.Error(fmt.Sprintf("get chainlink request template failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) chainLinkExpenseList(gin *gin.Context) {
	idStr := gin.Param("id")
	subscriptionId, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("subscription id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	requestName := gin.Query("requestName")
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
		logger.Error(fmt.Sprintf("user not found: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	data, err := h.chainLinkRequestService.ChainLinkExpenseList(subscriptionId, page, size, userId.(uint), requestName)
	if err != nil {
		logger.Error(fmt.Sprintf("query chain link expense list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) saveChainLinkRequestExec(gin *gin.Context) {
	createData := vo.ChainLinkRequestExecParam{}
	err := gin.BindJSON(&createData)
	if err != nil {
		logger.Error(fmt.Sprintf("create chainlink request exec failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	data, err := h.chainLinkRequestService.SaveChainLinkRequestExec(createData, userId.(uint))
	if err != nil {
		logger.Error(fmt.Sprintf("save request exec failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) getRequestId(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("request id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	data, err := h.chainLinkRequestService.GetRequestById(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) updateChainLinkRequestById(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error(fmt.Sprintf("subscription id question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	updateData := vo.ChainLinkExecParam{}
	err = gin.BindJSON(&updateData)
	if err != nil {
		logger.Error(fmt.Sprintf("body data question: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	userId, exists := gin.Get("userId")
	if !exists {
		logger.Error(fmt.Sprintf("user not found: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	err = h.chainLinkRequestService.UpdateChainLinkRequestById(int64(id), updateData, userId.(uint))
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) overview(gin *gin.Context) {
	userId, exists := gin.Get("userId")
	if !exists {
		Fail("do not have token", gin)
		return
	}
	// 获取路径参数
	network := gin.Param("network")
	appResp, err := h.chainLinkRequestService.Overview(userId.(uint), network)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(appResp, gin)

}
