package utils

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"hamster-paas/pkg/models/node"
	"testing"
	"time"
)

func TestSendEmail(t *testing.T) {
	SendEmail("sun1275596256@gmail.com", "0x49395e4a23265486c6a222a7f11e1bb158cc621c0278312f13976f50855b7369", "168880", "test", "")
}

func TestSendEmailForNodeCreate(t *testing.T) {
	rpcNode := node.RPCNode{
		Id:                1,
		Name:              "abing-test",
		UserId:            0,
		ChainProtocol:     "Ethereum",
		Status:            3,
		PublicIp:          "",
		Region:            "US East",
		LaunchTime:        sql.NullTime{Time: time.Now()},
		Resource:          "4C32GB 3000GB",
		ChainVersion:      "",
		NextPaymentDate:   sql.NullTime{Time: time.Now()},
		PaymentPerMonth:   decimal.Decimal{},
		RemainingSyncTime: "",
		CurrentHeight:     0,
		BlockTime:         "",
		HttpEndpoint:      "",
		WebsocketEndpoint: "",
		Created:           sql.NullTime{Time: time.Now()},
	}
	SendEmailForNodeCreate(rpcNode)
}
