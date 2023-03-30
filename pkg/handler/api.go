package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hamster-paas/pkg/utils/logger"
	"os"
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
	r.Use(h.handlerServer.Authorize())
	r.GET("/chains", chains)
	r.GET("/networks/:chain", networks)
	r.GET("/apps/:account", getApps)
	r.POST("/app", createApp)
	r.DELETE("/app/:account/:appId", deleteApp)

	// subscription
	r.GET("/subscription/overview", getSubscriptionOverview)

	return r.Run(fmt.Sprintf("0.0.0.0:%s", h.port))
}
