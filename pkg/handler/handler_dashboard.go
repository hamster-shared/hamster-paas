package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) dashboardAll(c *gin.Context) {
	_, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	all := h.chainlinkDashboardService.GetDashboardAll()
	Success(all, c)
}
