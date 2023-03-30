package handler

import (
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

func rpcGetChains(c *gin.Context) {
	chains, err := models.GetChains()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(chains, c)
}

func rpcGetNetworks(c *gin.Context) {
	// 路径参数
	chain, ok := c.Params.Get("chain")
	if !ok {
		Fail("invalid params", c)
		return
	}
	chainType, err := models.ParseChainType(chain)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	networks, err := models.GetNetworks(chainType)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(networks, c)
}

func rpcGetApps(c *gin.Context) {
	account, ok := c.Params.Get("account")
	if !ok {
		Fail("invalid params", c)
		return
	}
	page := c.Query("page")
	size := c.Query("size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		Fail("invalid params", c)
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		Fail("invalid params", c)
		return
	}
	var pagination models.Pagination
	pagination.Page = pageInt
	pagination.Size = sizeInt
	a, err := models.GetAccount(account)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	apps, p, err := a.GetApps(pagination)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	SuccessWithPagination(apps, p, c)
}

func rpcCreateApp(c *gin.Context) {
	var appParams vo.ApiRequestRpcCreateApp
	if err := c.ShouldBindJSON(&appParams); err != nil {
		Fail("invalid params", c)
		return
	}
	chain, err := models.ParseChainType(appParams.Chain)
	if err != nil {
		Fail("invalid params for chain", c)
		return
	}
	network, err := models.ParseNetworkType(appParams.Network)
	if err != nil {
		Fail("invalid params for network", c)
		return
	}
	a, err := models.GetAccount(appParams.Account)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	app, err := a.CreateApp(appParams.Name, appParams.Description, chain, network)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(app, c)
}

func rpcDeleteApp(c *gin.Context) {
	account, ok := c.Params.Get("account")
	if !ok {
		Fail("invalid params", c)
		return
	}
	appId, ok := c.Params.Get("appId")
	if !ok {
		Fail("invalid params", c)
		return
	}
	appIdInt, err := strconv.Atoi(appId)
	if err != nil {
		Fail("invalid params", c)
		return
	}
	a, err := models.GetAccount(account)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	_, err = a.GetApp(appIdInt)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = a.DeleteApp(appIdInt)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("", c)
}
