package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"
)

func (h *HandlerServer) RequestList(gin *gin.Context) {
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
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	data, err := h.chainLinkRequestService.RequestList(page, size, int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("get request list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) SaveChainLinkRequest(gin *gin.Context) {
	createData := vo.ChainLinkRequestParam{}
	err := gin.BindJSON(&createData)
	if err != nil {
		logger.Error(fmt.Sprintf("create chainlink request failed: %s", err.Error()))
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
	data := vo.ChainLinkRequest{
		Name:   createData.Name,
		Script: createData.Script,
		UserId: uint64(user.Id),
	}
	err = h.chainLinkRequestService.SaveChainLinkRequest(data)
	if err != nil {
		logger.Error(fmt.Sprintf("save chainlink request failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) UpdateChainLinkRequest(gin *gin.Context) {
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
	userAny, exists := gin.Get("user")
	if !exists {
		logger.Error(fmt.Sprintf("request list failed: %s", err.Error()))
		Fail("user information does not exist", gin)
		return
	}
	user, _ := userAny.(aline.User)
	data := vo.ChainLinkRequest{
		Name:   updateData.Name,
		Script: updateData.Script,
		UserId: uint64(user.Id),
	}
	err = h.chainLinkRequestService.UpdateChainLinkRequest(int64(id), data)
	if err != nil {
		logger.Error(fmt.Sprintf("update chainlink request failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) ChainLinkRequestTemplateList(gin *gin.Context) {
	data, err := h.chainLinkRequestService.ChainLinkRequestTemplateList()
	if err != nil {
		logger.Error(fmt.Sprintf("get chainlink request template failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) GetRequestTemplateScript(gin *gin.Context) {
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
