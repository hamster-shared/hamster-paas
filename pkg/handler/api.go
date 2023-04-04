package handler

import (
	"fmt"
	"hamster-paas/pkg/utils/logger"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	handlerServer HandlerServer
	port          string
}

func NewHttpService(handlerServer HandlerServer, port string) *HttpServer {
	return &HttpServer{
		handlerServer: handlerServer,
		port:          port,
	}
}

func (h *HttpServer) StartHttpServer() error {
	logger.Infof("start api server on port %s", h.port)
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()

	rpcApi := r.Group("/api/rpc")
	rpcApi.GET("/chains", rpcGetChains)
	rpcApi.GET("/networks/:chain", rpcGetNetworks)
	rpcApi.GET("/apps/:account", rpcGetApps)
	rpcApi.POST("/app", rpcCreateApp)
	rpcApi.DELETE("/app/:account/:appId", rpcDeleteApp)

	chainLinkApi := r.Group("/api/chainlink")
	chainLinkApi.Use(h.handlerServer.Authorize())
	//chain link request
	chainLinkApi.GET("/requests", h.handlerServer.requestList)
	chainLinkApi.POST("/request", h.handlerServer.saveChainLinkRequest)
	chainLinkApi.POST("/request/exec", h.handlerServer.saveChainLinkRequestExec)
	chainLinkApi.PUT("/request/:id", h.handlerServer.updateChainLinkRequest)
	chainLinkApi.GET("/request/templates", h.handlerServer.chainLinkRequestTemplateList)
	chainLinkApi.GET("/request/templates/:id", h.handlerServer.getRequestTemplateScript)
	// chain link subscription
	chainLinkApi.GET("/subscription/overview", h.handlerServer.getSubscriptionOverview)
	chainLinkApi.POST("/subscription/subscription", h.handlerServer.createSubscription)
	chainLinkApi.GET("/subscription/:id", h.handlerServer.subscriptionDetail)
	chainLinkApi.GET("/subscriptions", h.handlerServer.subscriptionList)
	chainLinkApi.POST("/subscription/:id/found", h.handlerServer.addFound)
	//// chain link consumer
	chainLinkApi.POST("//consumer/add-consumer", h.handlerServer.createConsumer)      //TODO...
	chainLinkApi.GET("/consumer/get-avail-consumer", h.handlerServer.getConsumerList) //TODO...
	chainLinkApi.GET("/subscription/:id/consumers", h.handlerServer.consumerList)
	chainLinkApi.DELETE("/subscription/:id/consumer/:consumerId", h.handlerServer.deleteConsumer)
	chainLinkApi.GET("/subscription/:id/expenses", h.handlerServer.chainLinkExpenseList)
	chainLinkApi.GET("/subscription/:id/deposits", h.handlerServer.depositList)
	chainLinkApi.GET("/dashboard/all", h.handlerServer.dashboardAll)
	chainLinkApi.GET("/dashboard/rpc", h.handlerServer.dashboardRpc)
	chainLinkApi.GET("/dashboard/oracle", h.handlerServer.dashboardOracle)

	return r.Run(fmt.Sprintf("0.0.0.0:%s", h.port))
}
