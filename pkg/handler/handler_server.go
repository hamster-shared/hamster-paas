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
	chainLinkDepositService      service.ChainLinkDepositService
}

func NewHandlerServer() *HandlerServer {
	handlerServer := HandlerServer{}
	chainLinkRequestService, err := application.GetBean[*service.ChainLinkRequestService]("chainLinkRequestService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainlink request service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink request service failed: %s", err.Error()))
	}
	handlerServer.chainLinkRequestService = *chainLinkRequestService

	chainLinkSubscriptionService, err := application.GetBean[*service.ChainLinkSubscriptionService]("chainLinkSubscriptionService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainlink subscription service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink subscription service failed: %s", err.Error()))
	}
	handlerServer.chainLinkSubscriptionService = *chainLinkSubscriptionService

	chainLinkConsumerService, err := application.GetBean[*service.ChainLinkConsumerService]("chainLinkConsumerService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainLink consumer service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainlink consumer service failed: %s", err.Error()))
	}
	handlerServer.chainLinkConsumerService = *chainLinkConsumerService
	chainLinkDepositService, err := application.GetBean[*service.ChainLinkDepositService]("chainLinkDepositService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainLink deposit service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainLink deposit service failed: %s", err.Error()))
	}
	handlerServer.chainLinkDepositService = *chainLinkDepositService

	return &handlerServer
}
