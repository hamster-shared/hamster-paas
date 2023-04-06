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
	chainlinkDashboardService    service.ChainLinkDashboardService
	chainlinkPoolService         service.PoolService
	rpcService                   service.RpcService
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
	chainLinkDashboardService, err := application.GetBean[*service.ChainLinkDashboardService]("chainLinkDashboardService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get chainLink dashboard service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get chainLink dashboard service failed: %s", err.Error()))
	}
	handlerServer.chainlinkDashboardService = *chainLinkDashboardService
	rpcService, err := application.GetBean[*service.RpcService]("rpcService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get rpc service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get rpc service failed: %s", err.Error()))
	}
	handlerServer.rpcService = *rpcService
	poolService, err := application.GetBean[*service.PoolService]("chainLinkPoolService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get pool service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get pool service failed: %s", err.Error()))
	}
	handlerServer.chainlinkPoolService = *poolService
	return &handlerServer
}
