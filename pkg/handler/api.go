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
	rpcApi.Use(h.handlerServer.Authorize())
	rpcApi.GET("/chains", h.handlerServer.rpcGetChains)
	rpcApi.GET("/networks/:chain", h.handlerServer.rpcGetNetworks)
	rpcApi.GET("/overview", h.handlerServer.rpcOverview)
	rpcApi.GET("/mynetwork", h.handlerServer.rpcGetMyNetwork)
	rpcApi.GET("/chain/:chain", h.handlerServer.rpcChainDetail)
	rpcApi.GET("/request-log/:appKey", h.handlerServer.rpcRequestLog)

	middleWare := r.Group("/api/middleware")
	middleWare.GET("/rpc", h.handlerServer.middlewareRpc)

	chainLinkApi := r.Group("/api/chainlink")
	chainLinkApi.Use(h.handlerServer.Authorize())
	//chain link request
	chainLinkApi.GET("/request/overview/:network", h.handlerServer.overview)
	chainLinkApi.GET("/requests", h.handlerServer.requestList)
	chainLinkApi.POST("/request", h.handlerServer.saveChainLinkRequest)
	chainLinkApi.POST("/request/exec", h.handlerServer.saveChainLinkRequestExec)
	chainLinkApi.PUT("/request/exec/:id", h.handlerServer.updateChainLinkRequestById)
	chainLinkApi.PUT("/request/:id", h.handlerServer.updateChainLinkRequest)
	chainLinkApi.GET("/request/templates", h.handlerServer.chainLinkRequestTemplateList)
	chainLinkApi.GET("/request/templates/:id", h.handlerServer.getRequestTemplateScript)
	// chain link subscription
	chainLinkApi.GET("/subscription/overview", h.handlerServer.getSubscriptionOverview)
	chainLinkApi.POST("/subscription/subscription", h.handlerServer.createSubscription)
	chainLinkApi.PUT("/subscription/subscription-status", h.handlerServer.changeSubscriptionStatus)
	chainLinkApi.GET("/subscription/:id", h.handlerServer.subscriptionDetail)
	chainLinkApi.GET("/subscription/valid-subscription", h.handlerServer.getValidSubscription)
	chainLinkApi.GET("/subscriptions", h.handlerServer.subscriptionList)
	chainLinkApi.POST("/subscription/:id/fund", h.handlerServer.addFound)
	chainLinkApi.PUT("/subscription/fund-status", h.handlerServer.changeFoundStatus)
	//// chain link consumer
	chainLinkApi.POST("/consumer", h.handlerServer.createConsumer)
	chainLinkApi.PUT("/consumer/consumer-status", h.handlerServer.changeConsumerStatus)
	chainLinkApi.GET("/consumer/:id/hamster-consumer", h.handlerServer.getHamsterConsumerList)
	chainLinkApi.GET("/consumer/projects", h.handlerServer.getProjectList)
	chainLinkApi.GET("/subscription/:id/consumers", h.handlerServer.consumerList)
	chainLinkApi.GET("/subscription/:id/consumer-address", h.handlerServer.consumerAddressList)
	chainLinkApi.DELETE("/subscription/:id/consumer/:consumerId", h.handlerServer.deleteConsumer)
	chainLinkApi.GET("/subscription/:id/expenses", h.handlerServer.chainLinkExpenseList)
	chainLinkApi.GET("/subscription/:id/deposits", h.handlerServer.depositList)
	chainLinkApi.GET("/dashboard/all", h.handlerServer.dashboardAll)

	return r.Run(fmt.Sprintf("0.0.0.0:%s", h.port))
}
