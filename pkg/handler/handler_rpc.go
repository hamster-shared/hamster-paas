package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/models"
	"log"
	"strconv"
)

func chains(c *gin.Context) {
	chains, err := models.GetChains()
	if err != nil {
		Fail(err.Error(), c)
		return
	}
	Success(chains, c)
}

func networks(c *gin.Context) {
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

func getApps(c *gin.Context) {
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

func createApp(c *gin.Context) {
	var appParams models.ApiRequestCreateApp
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

func deleteApp(c *gin.Context) {
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

func getSubscriptionOverview(c *gin.Context) {
	userId := c.Query("userid")
	network := c.Query("network")

	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		panic(err)
	}

	type overview struct {
		TotalSubscription int
		TotalConsumers    int
		TotalBalance      float64
	}
	var ov overview

	sql := "select COUNT(*) as total_subscription, SUM(consumers) as total_consumers, SUM(balance) as total_balance from t_cl_subscription where user_id = ? AND chain = ?"
	err = db.Raw(sql, userId, network).Scan(&ov).Error
	if err != nil {
		log.Println(err)
	}

	Success(ov, c)
}
