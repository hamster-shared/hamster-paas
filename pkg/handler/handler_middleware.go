package handler

import (
	"fmt"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) middlewareRpc(c *gin.Context) {
	userAny, exit := c.Get("user")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", c)
		return
	}
	loginType, exit := c.Get("loginType")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", c)
		return
	}
	var err error
	var result []vo.MiddleWareRpcZan
	if loginType == consts.GitHub {
		result, err = h.middleWareService.MiddleWareRpc(fmt.Sprintf("%d", userAny.(aline.User).Id))
	} else if loginType == consts.Metamask {
		result, err = h.middleWareService.MiddleWareRpc(fmt.Sprintf("%d", userAny.(aline.UserWallet).UserId))
	}

	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(result, c)
}

func (h *HandlerServer) serviceIsActive(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	serviceName := c.Param("serviceName")
	response := h.rpcService.IsActive(user.(aline.User), serviceName)
	Success(response, c)
}

func (h *HandlerServer) activeService(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	serviceName := c.Param("serviceName")
	type ApiRequest struct {
		Chain   string `json:"chain"`
		Network string `json:"network"`
	}
	var apiRequest ApiRequest
	if strings.ToLower(serviceName) == "rpc" {
		err := c.ShouldBindJSON(&apiRequest)
		if err != nil {
			Fail(fmt.Sprintf("invalid params, need 'chain' and 'network', err: %s", err), c)
			return
		}
	}
	msg, err := h.rpcService.ActiveService(user.(aline.User), serviceName, apiRequest.Chain, apiRequest.Network)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(msg, c)
}
