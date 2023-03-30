package handler

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/utils/logger"
)

type HandlerServer struct {
	chainLinkRequestService service.ChainLinkRequestService
}

func NewHandlerServer() *HandlerServer {
	handlerServer := HandlerServer{}
	chainLinkRequestService, err := application.GetBean[*service.ChainLinkRequestService]("chainLinkRequestService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
	}
	handlerServer.chainLinkRequestService = *chainLinkRequestService
	return &handlerServer
}
