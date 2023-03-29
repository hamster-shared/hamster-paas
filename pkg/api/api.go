package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/handler"
	"hamster-paas/pkg/logger"
	"hamster-paas/pkg/models"
	"log"
	"os"
	"strconv"
)

type HttpServer struct {
	handlerServer handler.HandlerServer
	port          string
}

func NewHttpService(handlerServer handler.HandlerServer, port string) *HttpServer {
	return &HttpServer{
		handlerServer: handlerServer,
		port:          port,
	}
}

func (h *HttpServer) StartHttpServer() {
	logger.Infof("start api server on port %s", h.port)
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.Use()
	r.GET("/chains", chains)
	r.GET("/networks/:chain", networks)
	r.GET("/apps/:account", getApps)
	r.POST("/app", createApp)
	r.DELETE("/app/:account/:appId", deleteApp)

	// subscription
	r.GET("/subscription/overview", getSubscriptionOverview)

	r.Run(fmt.Sprintf("0.0.0.0:%s", h.port))
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

	Success(c, ov)
}
