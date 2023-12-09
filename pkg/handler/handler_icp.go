package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) DfxVersion(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	ver, err := h.icpService.GetDfxVersion()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(ver, c)
}
