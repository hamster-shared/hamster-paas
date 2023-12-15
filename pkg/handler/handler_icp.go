package handler

import (
	"hamster-paas/pkg/models/vo"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// DfxVersion 获取 DFX 版本
func (h *HandlerServer) DfxVersion(c *gin.Context) {
	version, err := h.icpService.GetDfxVersion()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	ver := strings.Fields(version)[1]
	Success(ver, c)
}

const (
	USERID = uint(65406422)
)

// IcpAccountBreif 获取账户概览
func (h *HandlerServer) IcpAccountBreif(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetAccountBrief(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpAccountOverview 获取账户概览
func (h *HandlerServer) IcpAccountOverview(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetAccountOverview(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// page, size
// IcpCanisterPage 获取指定页面的 canister
func (h *HandlerServer) IcpCanisterPage(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	data, err := h.icpService.GetCanisterPage(userId.(uint), page, size)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpCanisterOverview 获取指定 canister 的概览
func (h *HandlerServer) IcpCanisterOverview(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	canisterId := c.Param("id")
	if canisterId == "" {
		Fail("canister id is empty", c)
		return
	}
	data, err := h.icpService.GetCanisterOverview(userId.(uint), canisterId)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpControllerPage 获取指定 canister 的 controller
func (h *HandlerServer) IcpControllerPage(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	pageStr := c.DefaultQuery("id", "1")
	sizeStr := c.DefaultQuery("size", "5")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	canisterId := c.Param("id")
	if canisterId == "" {
		Fail("canister id is empty", c)
		return
	}
	data, err := h.icpService.GetContollerPage(userId.(uint), canisterId, page, size)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpConsumptionPage 获取指定 canister 的消费
func (h *HandlerServer) IcpConsumptionPage(c *gin.Context) {
	pageStr := c.DefaultQuery("id", "1")
	sizeStr := c.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	canisterId := c.Param("id")
	if canisterId == "" {
		Fail("canister id is empty", c)
		return
	}
	data, err := h.icpService.GetConsumptionPage(canisterId, page, size)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpGetAccount 获取钱包账户启动
func (h *HandlerServer) IcpGetAccount(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetAccountFlag(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpAccountInfo 获取钱包账户信息
func (h *HandlerServer) IcpAccountInfo(c *gin.Context) {
	userId, exists := c.Get("userId")
	icpTest := os.Getenv("IPC_TEST")
	if icpTest == "true" {
		userId = USERID
		exists = true
	}
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetAccountInfo(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// POST IcpCreateIdentity 创建钱包账户信息
func (h *HandlerServer) IcpCreateIdentity(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.CreateIdentity(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpGetWalletCycles 获取钱包 Gas
func (h *HandlerServer) IcpGetWalletCycles(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.GetWalletInfo(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

// IcpRechargeWallet 购买钱包 Gas
func (h *HandlerServer) IcpRechargeWallet(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	data, err := h.icpService.RechargeWallet(userId.(uint))
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

func (h *HandlerServer) IcpAddCanister(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var addCanisterParam vo.CreateCanisterParam
	err := c.BindJSON(&addCanisterParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.AddCanister(userId.(uint), addCanisterParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}

func (h *HandlerServer) IcpDelCanister(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var delCanisterParam vo.DeleteCanisterParam
	err := c.BindJSON(&delCanisterParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.DelCanister(userId.(uint), delCanisterParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}

func (h *HandlerServer) IcpRechargeCanister(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var rechargeCanisterParam vo.RechargeCanisterParam
	err := c.BindJSON(&rechargeCanisterParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	data, err := h.icpService.RechargeCanister(userId.(uint), rechargeCanisterParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(data, c)
}

func (h *HandlerServer) IcpAddController(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var addControllerParam vo.AddControllerParam
	err := c.BindJSON(&addControllerParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.AddController(userId.(uint), addControllerParam)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}

func (h *HandlerServer) IcpDelController(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var param vo.DelControllerParam
	err := c.BindJSON(&param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.DelController(userId.(uint), param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}

func (h *HandlerServer) IcpChangeStatus(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var param vo.ChangeStatusParam
	err := c.BindJSON(&param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.ChangeCanisterStatus(userId.(uint), param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}

func (h *HandlerServer) IcpDeleteCanister(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var param vo.ChangeStatusParam
	err := c.BindJSON(&param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.ChangeCanisterStatus(userId.(uint), param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}

func (h *HandlerServer) IcpInstallDapp(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		Fail("do not have token", c)
		return
	}
	var param vo.InstallDappParam
	err := c.BindJSON(&param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	err = h.icpService.InstallDapp(userId.(uint), param)
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success("SUCCESS", c)
}
