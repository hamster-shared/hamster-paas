package api

import (
	"fmt"
	"hamster-paas/pkg/models"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.GET("/chains", chains)
	r.GET("/networks/:chain", networks)
	r.GET("/apps/:account", getApps)
	r.POST("/app", createApp)
	r.DELETE("/app/:account/:appId", deleteApp)

	r.GET("/subscription/overview", subscriptionOverview)

	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func chains(c *gin.Context) {
	chains, err := models.GetChains()
	if err != nil {
		Fail(c, err.Error())
		return
	}
	Success(c, chains)
}

func networks(c *gin.Context) {
	// 路径参数
	chain, ok := c.Params.Get("chain")
	if !ok {
		Fail(c, "invalid params")
		return
	}
	chainType, err := models.ParseChainType(chain)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	networks, err := models.GetNetworks(chainType)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	Success(c, networks)
}

func getApps(c *gin.Context) {
	account, ok := c.Params.Get("account")
	if !ok {
		Fail(c, "invalid params")
		return
	}
	page := c.Query("page")
	size := c.Query("size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		Fail(c, "invalid params")
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		Fail(c, "invalid params")
		return
	}
	var pagination models.Pagination
	pagination.Page = pageInt
	pagination.Size = sizeInt
	a, err := models.GetAccount(account)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	apps, p, err := a.GetApps(pagination)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	SuccessWithPagination(c, apps, p)
}

func createApp(c *gin.Context) {
	var appParams models.ApiRequestCreateApp
	if err := c.ShouldBindJSON(&appParams); err != nil {
		Fail(c, "invalid params")
		return
	}
	chain, err := models.ParseChainType(appParams.Chain)
	if err != nil {
		Fail(c, "invalid params for chain")
		return
	}
	network, err := models.ParseNetworkType(appParams.Network)
	if err != nil {
		Fail(c, "invalid params for network")
		return
	}
	a, err := models.GetAccount(appParams.Account)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	app, err := a.CreateApp(appParams.Name, appParams.Description, chain, network)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	Success(c, app)
}

func deleteApp(c *gin.Context) {
	account, ok := c.Params.Get("account")
	if !ok {
		Fail(c, "invalid params")
		return
	}
	appId, ok := c.Params.Get("appId")
	if !ok {
		Fail(c, "invalid params")
		return
	}
	appIdInt, err := strconv.Atoi(appId)
	if err != nil {
		Fail(c, "invalid params")
		return
	}
	a, err := models.GetAccount(account)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	_, err = a.GetApp(appIdInt)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	err = a.DeleteApp(appIdInt)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	Success(c, nil)
}

func subscriptionOverview(c *gin.Context) {

}
