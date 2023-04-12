package handler

import (
	"hamster-paas/pkg/rpc/aline"

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
	ok = h.rpcService.IsActive(user.(aline.User), serviceName)
	Success(ok, c)
}

func (h *HandlerServer) activeService(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	serviceName := c.Param("serviceName")
	msg := h.rpcService.ActiveService(user.(aline.User), serviceName)
	Success(msg, c)
}
