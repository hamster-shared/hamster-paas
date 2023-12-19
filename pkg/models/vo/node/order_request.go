package node

import (
	"github.com/shopspring/decimal"
	"hamster-paas/pkg/models/order"
)

type LaunchOrderParam struct {
	Protocol     string          `json:"protocol"`
	Region       string          `json:"region"`
	ResourceType string          `json:"resourceType"`
	NodeResource string          `json:"nodeResource"`
	BuyTime      int             `json:"buyTime"`
	NodeName     string          `json:"nodeName"`
	Amount       decimal.Decimal `json:"amount"`
	PayType      order.PayType   `json:"payType"`
}
