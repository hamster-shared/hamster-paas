package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"
)

func (h *HandlerServer) depositList(gin *gin.Context) {
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
	data, err := h.chainLinkDepositService.DepositList(page, size, int64(user.Id))
	if err != nil {
		logger.Error(fmt.Sprintf("query chain link deposit list failed: %s", err.Error()))
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
