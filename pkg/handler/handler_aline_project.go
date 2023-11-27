package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
)

func (h *HandlerServer) getProjectList(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	chain := c.Query("chain")
	network := c.Query("network")
	if chain == "" || network == "" {
		Fail("param invalid", c)
		return
	}
	network = models.GetAlineNetwork(chain, network)

	projectService, err := application.GetBean[*aline.ProjectService]("projectService")
	if err != nil {
		Fail("get project service error", c)
		return
	}
	projectIdAndNameList := projectService.GetProjectByUserId(userId.(uint), network)
	Success(projectIdAndNameList, c)
}
