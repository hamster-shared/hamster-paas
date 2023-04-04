package handler

import (
	"hamster-paas/pkg/rpc/aline"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) dashboardAll(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	all := h.chainlinkDashboardService.GetDashboardAll(user.(aline.User))
	Success(all, c)
}
