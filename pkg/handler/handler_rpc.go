package handler

import (
	"fmt"
	"hamster-paas/pkg/models"
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
	network := c.Param("network")
	if network == "" {
		Fail(fmt.Sprintf("invalid params, network: %s", network), c)
		return
	}
	appResp, err := h.rpcService.Overview(user.(aline.User), network)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(appResp, c)
}

func (h *HandlerServer) rpcGetMyNetwork(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	page := c.Query("page")
	size := c.Query("size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		sizeInt = 10
	}
	p := &models.Pagination{
		Page: pageInt,
		Size: sizeInt,
	}
	appResp, p, err := h.rpcService.GetMyNetwork(user.(aline.User), p)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	SuccessWithPagination(appResp, *p, c)
}

func (h *HandlerServer) rpcGetChains(c *gin.Context) {
	chains, err := h.rpcService.GetChains()
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

func (h *HandlerServer) rpcChainDetail(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	chain, ok := c.Params.Get("chain")
	if !ok {
		Fail("invalid params", c)
		return
	}
	chainDetail, err := h.rpcService.ChainDetail(user.(aline.User), chain)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(chainDetail, c)
}

func (h *HandlerServer) rpcRequestLog(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		Fail("do not have token", c)
		return
	}
	appKey, ok := c.Params.Get("appKey")
	if !ok {
		Fail("invalid params", c)
		return
	}
	var page, size string
	page = c.Query("page")
	if page == "" {
		page = "1"
	}
	size = c.Query("size")
	if size == "" {
		size = "10"
	}
	requestLog, p, err := h.rpcService.AppRequestLog(user.(aline.User), appKey, page, size)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	SuccessWithPagination(requestLog, *p, c)
}
