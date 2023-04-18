package handler

import (
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
)

func (h *HandlerServer) getProjectList(c *gin.Context) {
	userAny, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	user := userAny.(aline.User)
	chain := c.Query("chain")
	network := c.Query("network")
	if chain == "" || network == "" {
		Fail("param invalid", c)
		return
	}
	network, err := models.GetAlineNetwork(chain, network)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	projectService, err := application.GetBean[*aline.ProjectService]("projectService")
	if err != nil {
		Fail("get project service error", c)
		return
	}
	projectIdAndNameList := projectService.GetProjectByUserId(user.Id, network)
	Success(projectIdAndNameList, c)
}
