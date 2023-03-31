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
	chainLinkApi.GET("/requests", h.handlerServer.RequestList)
	chainLinkApi.POST("/request", h.handlerServer.SaveChainLinkRequest)
	chainLinkApi.PUT("/request/:id", h.handlerServer.UpdateChainLinkRequest)
	chainLinkApi.GET("/request/templates", h.handlerServer.ChainLinkRequestTemplateList)
	chainLinkApi.GET("/request/templates/:id", h.handlerServer.GetRequestTemplateScript)
	// chain link subscription
	chainLinkApi.GET("/subscription/overview", h.handlerServer.getSubscriptionOverview)
	chainLinkApi.GET("/subscription/sina", h.handlerServer.getSINA)
	chainLinkApi.POST("/subscription/create-subscription", h.handlerServer.createSubscription)
	// chain link consumer
	chainLinkApi.POST("/consumer/add-consumer", h.handlerServer.createConsumer)
	chainLinkApi.GET("/consumer/get-avail-consumer", h.handlerServer.getConsumerList)
	// chain link found
	chainLinkApi.PUT("/found/add-found", h.handlerServer.addFound)

	return r.Run(fmt.Sprintf("0.0.0.0:%s", h.port))
}
