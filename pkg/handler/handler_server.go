package handler

import (
	"fmt"
	"hamster-paas/pkg/application"
	"hamster-paas/pkg/service"
	service2 "hamster-paas/pkg/service/node"
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
	middleWareService            service.MiddleWareService
	nodeService                  service2.NodeService
	orderService                 service2.OrderService
	resourceStandardService      service2.ResourceStandardService
	zanService                   service.ZanService
	icpService                   service.IcpService
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
	middleWareService, err := application.GetBean[*service.MiddleWareService]("middleWareService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get middleware service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get middleware service failed: %s", err.Error()))
	}
	handlerServer.middleWareService = *middleWareService

	nodeService, err := application.GetBean[*service2.NodeService]("nodeService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get node service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get node service failed: %s", err.Error()))
	}
	handlerServer.nodeService = *nodeService
	orderService, err := application.GetBean[*service2.OrderService]("orderService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get order service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get order service failed: %s", err.Error()))
	}
	handlerServer.orderService = *orderService

	resourceStandardService, err := application.GetBean[*service2.ResourceStandardService]("resourceStandardService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get resource standard service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get resource stanard service failed: %s", err.Error()))
	}
	handlerServer.resourceStandardService = *resourceStandardService

	zanService, err := application.GetBean[*service.ZanService]("zanService")
	if err != nil {
		logger.Error(fmt.Sprintf("application get zan service failed: %s", err.Error()))
		panic(fmt.Sprintf("application get zan service failed: %s", err.Error()))
	}

	handlerServer.zanService = *zanService

	return &handlerServer
}
