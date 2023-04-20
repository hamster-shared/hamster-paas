package handler

import (
	"fmt"
	"hamster-paas/pkg/rpc/aline"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) middlewareRpc(c *gin.Context) {
	result, err := h.middleWareService.MiddleWareRpc()
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
