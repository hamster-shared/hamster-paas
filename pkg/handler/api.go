package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketIo "github.com/googollee/go-socket.io"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hamster-paas/docs"
	"hamster-paas/pkg/utils/logger"
	"os"
)

type HttpServer struct {
	handlerServer  HandlerServer
	port           string
	socketIoServer *socketIo.Server
}

func NewHttpService(handlerServer HandlerServer, port string, socketIoServer *socketIo.Server) *HttpServer {
	return &HttpServer{
		handlerServer:  handlerServer,
		port:           port,
		socketIoServer: socketIoServer,
	}
}

func (h *HttpServer) StartHttpServer() error {
	logger.Infof("start api server on port %s", h.port)
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//socket
	go func() {
		if err := h.socketIoServer.Serve(); err != nil {
			logger.Errorf("socketIo listen error: %s\n", err)
		}
	}()
	defer h.socketIoServer.Close()
	socketApi := r.Group("/socket.io")
	socketApi.GET("/*any", gin.WrapH(h.socketIoServer))
	socketApi.POST("/*any", gin.WrapH(h.socketIoServer))

	rpcApi := r.Group("/api/rpc")
	rpcApi.Use(h.handlerServer.Authorize())
	rpcApi.GET("/chains", h.handlerServer.rpcGetChains)
	rpcApi.GET("/networks/:chain", h.handlerServer.rpcGetNetworks)
	rpcApi.GET("/overview/:network", h.handlerServer.rpcOverview)
	rpcApi.GET("/mynetwork", h.handlerServer.rpcGetMyNetwork)
	rpcApi.GET("/chain/:chain", h.handlerServer.rpcChainDetail)
	rpcApi.GET("/request-log/:appKey", h.handlerServer.rpcRequestLog)

	middleWare := r.Group("/api/middleware")
	middleWare.Use(h.handlerServer.Authorize())
	middleWare.GET("/rpc", h.handlerServer.middlewareRpc)
	middleWare.GET("/is-active/:serviceName", h.handlerServer.serviceIsActive)
	middleWare.POST("/active/:serviceName", h.handlerServer.activeService)

	chainLinkApi := r.Group("/api/chainlink")
	chainLinkApi.Use(h.handlerServer.Authorize())
	//chain link request
	chainLinkApi.GET("/request/overview/:network", h.handlerServer.overview)
	chainLinkApi.GET("/requests", h.handlerServer.requestList)
	chainLinkApi.POST("/request", h.handlerServer.saveChainLinkRequest)
	chainLinkApi.POST("/request/exec", h.handlerServer.saveChainLinkRequestExec)
	chainLinkApi.GET("/request/:id", h.handlerServer.getRequestId)
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
	chainLinkApi.GET("/subscription/balance", h.handlerServer.getSubscriptionBalanceAll)
	chainLinkApi.GET("/subscription/:id/balance", h.handlerServer.getSubscriptionBalanceById)
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

	//node api
	nodeApi := r.Group("/api/node")
	nodeApi.Use(h.handlerServer.Authorize())
	nodeApi.GET("/list", h.handlerServer.nodeList)
	nodeApi.GET("/statistics-info", h.handlerServer.nodeStatisticsInfo)
	nodeApi.GET("/:id", h.handlerServer.nodeDetail)
	nodeApi.POST("/order/update", h.handlerServer.updateNode)
	nodeApi.POST("/order/launch", h.handlerServer.launchOrder)
	nodeApi.GET("/order/list", h.handlerServer.orderList)
	nodeApi.GET("/order/:id", h.handlerServer.payOrderDetail)
	nodeApi.PUT("/order/:id/cancel", h.handlerServer.cancelOrder)
	nodeApi.GET("/resource-standard/:protocol", h.handlerServer.queryResourceStandard)

	zanApi := r.Group("/api/v2/zan")
	zanApi.Use(h.handlerServer.Authorize())
	zanApi.GET("/account/authed", h.handlerServer.ZanAuthed)
	zanApi.GET("/account/auth_url", h.handlerServer.ZanGetAuthUrl)
	zanApi.POST("/account/access_token", h.handlerServer.ZanExchangeAccessToken)
	zanApi.POST("/node-service/api-keys", h.handlerServer.ZanCreateApiKey)
	zanApi.GET("/node-service/api-keys/list", h.handlerServer.ZanApiKeyPage)
	zanApi.GET("/node-service/api-keys/detail", h.handlerServer.ZanApiKeyDetail)
	zanApi.GET("/node-service/api-keys/stats/credit-cost", h.handlerServer.ZanApiKeyCreditCost)
	zanApi.GET("/node-service/api-keys/stats/requests", h.handlerServer.ZanApiKeyRequestStats)
	zanApi.GET("/node-service/api-keys/stats/requests-activity", h.handlerServer.ZanApiKeyRequestActivityStats)
	zanApi.GET("/node-service/api-keys/stats/requests-origin", h.handlerServer.ZanApiKeyRequestsOriginStats)
	zanApi.GET("/node-service/ecosystems/digest", h.handlerServer.ZanEcosystemsDigest)
	zanApi.GET("/node-service/plan", h.handlerServer.ZanPlan)
	return r.Run(fmt.Sprintf("0.0.0.0:%s", h.port))
}
