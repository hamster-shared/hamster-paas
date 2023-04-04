package handler

import (
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/rpc/aline"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) rpcOverview(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	appResp, err := h.rpcService.Overview(user.(aline.User))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(appResp, c)
}

func (h *HandlerServer) rpcGetChains(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	chains, err := h.rpcService.GetChains(user.(aline.User))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(chains, c)
}

func (h *HandlerServer) rpcGetNetworks(c *gin.Context) {
	chain, ok := c.Params.Get("chain")
	if !ok {
		Fail("invalid params", c)
		return
	}
	networks, err := h.rpcService.GetNetworks(chain)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(networks, c)
}

// func (h *HandlerServer) rpcGetApps(c *gin.Context) {
// 	user, ok := c.Get("user")
// 	if !ok {
// 		Fail("do not have token", c)
// 		return
// 	}
// 	chains, err := h.rpcService.GetApps(user.(aline.User))
// 	if err != nil {
// 		Fail(err.Error(), c)
// 		return
// 	}
// 	Success(chains, c)

// }

func rpcGetApps(c *gin.Context) {
	account, ok := c.Params.Get("account")
	if !ok {
		Fail("invalid params", c)
		return
	}
	a, err := models.GetRpcAccount(account)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	apps, err := a.GetApps()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(apps, c)
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
	a, err := models.GetRpcAccount(appParams.Account)
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
	a, err := models.GetRpcAccount(account)
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
