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

func (h *HandlerServer) dashboardRpc(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	rpc := h.chainlinkDashboardService.GetDashboardRpc(user.(aline.User))
	Success(rpc, c)
}

func (h *HandlerServer) dashboardOracle(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	oracle := h.chainlinkDashboardService.GetDashboardOracle(user.(aline.User))
	Success(oracle, c)
}
