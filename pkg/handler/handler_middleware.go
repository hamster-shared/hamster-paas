package handler

import (
	"fmt"
	"hamster-paas/pkg/models/vo"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) middlewareRpc(c *gin.Context) {
	userId, exit := c.Get("userId")
	if !exit {
		Failed(http.StatusUnauthorized, "access not authorized", c)
		return
	}
	var err error
	var result []vo.MiddleWareRpcZan
	result, err = h.middleWareService.MiddleWareRpc(fmt.Sprintf("%d", userId.(uint)))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(result, c)
}

func (h *HandlerServer) serviceIsActive(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		Fail("do not have token", c)
		return
	}
	serviceName := c.Param("serviceName")
	response := h.rpcService.IsActive(int(userId.(uint)), serviceName)
	Success(response, c)
}

func (h *HandlerServer) activeService(c *gin.Context) {
	userId, ok := c.Get("userId")
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
	msg, err := h.rpcService.ActiveService(userId.(uint), serviceName, apiRequest.Chain, apiRequest.Network)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(msg, c)
}
