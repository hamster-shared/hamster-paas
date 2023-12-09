package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) DfxVersion(c *gin.Context) {
	version, err := h.icpService.GetDfxVersion()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	ver := strings.Fields(version)[1]
	Success(ver, c)
}

func (h *HandlerServer) IcpAccountBreif(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetAccountBrief(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

func (h *HandlerServer) IcpAccountOverview(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetAccountOverview(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}
