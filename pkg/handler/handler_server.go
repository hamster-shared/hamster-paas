package handler

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/service"
	"hamster-paas/pkg/utils/logger"
)

type HandlerServer struct {
	chainLinkRequestService      service.ChainLinkRequestService
	chainLinkSubscriptionService service.ChainLinkSubscriptionService
	chainLinkConsumerService     service.ChainLinkConsumerService
}

func NewHandlerServer() *HandlerServer {
	handlerServer := HandlerServer{}
	chainLinkRequestService, err := application.GetBean[*service.ChainLinkRequestService]("chainLinkRequestService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
	}
	handlerServer.chainLinkRequestService = *chainLinkRequestService

	chainLinkSubscriptionService, err := application.GetBean[*service.ChainLinkSubscriptionService]("chainLinkSubscriptionService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
	}
	handlerServer.chainLinkSubscriptionService = *chainLinkSubscriptionService

	chainLinkConsumerService, err := application.GetBean[*service.ChainLinkConsumerService]("chainLinkConsumerService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink service failed: %s", err.Error()))
	}
	handlerServer.chainLinkConsumerService = *chainLinkConsumerService

	return &handlerServer
}
