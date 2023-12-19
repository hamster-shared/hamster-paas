package node

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	modelsNode "hamster-paas/pkg/models/node"
	"hamster-paas/pkg/models/order"
	"hamster-paas/pkg/models/vo/node"
	"hamster-paas/pkg/service/contract"
	"hamster-paas/pkg/utils/logger"
	"math/big"
	"os"
	"time"
)

type OrderService struct {
	db            *gorm.DB
	snowFlakeNode *snowflake.Node
}

func NewOrderService(db *gorm.DB) *OrderService {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Errorf("get snowflake failed: %s", err)
		panic(err)
	}
	return &OrderService{
		db:            db,
		snowFlakeNode: node,
	}
}

func (o *OrderService) LaunchOrder(userId uint, launchData node.LaunchOrderParam) (uint, error) {
	var queryOrder order.Order
	err := o.db.Model(order.Order{}).Where("user_id = ? and status = ?", userId, order.PaymentPending).First(&queryOrder).Error
	if err == nil {
		return 0, fmt.Errorf("there are unpaid orders")
	}
	var orderNodeData order.OrderNode
	err = o.db.Model(order.OrderNode{}).Where("user_id = ? and node_name = ?", userId, launchData.NodeName).First(&orderNodeData).Error
	if err == nil {
		return 0, fmt.Errorf("node name: %s is areadly exists", launchData.NodeName)
	}
	var changeAccount modelsNode.UserChargeAccount
	var orderData order.Order
	err = o.db.Model(modelsNode.UserChargeAccount{}).Where("user_id = ?", userId).First(&changeAccount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			address, privateKey, err := GenerateEthAddress()
			if err != nil {
				logger.Errorf("generate address failed: %s", err)
				return 0, err
			}
			changeAccount.UserId = int(userId)
			changeAccount.Seed = privateKey
			changeAccount.Address = address
			err = o.db.Create(&changeAccount).Error
			if err != nil {
				logger.Errorf("save change account failed: %s", err)
				return 0, fmt.Errorf("save change account failed: %s", err)
			}
		} else {
			logger.Errorf("query change account failed: %s", err)
			return 0, err
		}
	}
	data, err := AccountBalance(changeAccount.Address)
	if err != nil {
		logger.Errorf("get balance failed: %s", err)
		return 0, err
	}
	balance, err := decimal.NewFromString(data)
	if err != nil {
		logger.Errorf("balance to decimal failed: %s", err)
		return 0, err
	}
	orderId := o.snowFlakeNode.Generate().String()
	orderData.OrderId = orderId
	orderData.OrderTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	orderData.AddressInitBalance = sql.NullString{
		String: balance.String(),
		Valid:  true,
	}
	orderData.OrderType = order.NodeService
	orderData.UserId = uint(userId)
	orderData.Status = order.PaymentPending
	orderData.Chain = modelsNode.ChainProtocol(launchData.Protocol)
	orderData.Amount = sql.NullString{
		String: launchData.Amount.String(),
		Valid:  true,
	}
	orderData.BuyTime = launchData.BuyTime
	orderData.ResourceType = launchData.ResourceType
	orderData.ReceiveAddress = changeAccount.Address
	orderData.NodeName = launchData.NodeName
	orderData.PayType = launchData.PayType
	err = o.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Save(&orderData).Error
		if err != nil {
			return err
		}
		var orderNodeData order.OrderNode
		orderNodeData.OrderId = orderData.Id
		orderNodeData.NodeName = launchData.NodeName
		orderNodeData.Protocol = launchData.Protocol
		orderNodeData.UserId = uint(userId)
		orderNodeData.Resource = launchData.NodeResource
		orderNodeData.Region = launchData.Region
		orderNodeData.CreateTime = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		err = tx.Save(&orderNodeData).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("launch order failed: %s", err)
		return 0, err
	}
	return orderData.Id, nil
}

// query order list
func (o *OrderService) OrderList(startDateStr, endDateStr, query string, userId uint, page, size int) (node.OrderPage, error) {
	var orderList []order.Order
	var resultData []node.OrderVo
	var pageData node.OrderPage
	var total int64
	tx := o.db.Model(order.Order{}).Where("user_id = ?", userId)
	if startDateStr != "" && endDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			logger.Errorf("start date parse failed: %s", err)
			return pageData, err
		}
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			logger.Errorf("end date parse failed: %s", err)
		}
		tx = tx.Where("order_time between ? and ?", startDate, endDate)
	}
	if query != "" {
		tx = tx.Where("order_id like ?", "%"+query+"%")
	}
	err := tx.Order("order_time DESC").Offset((page - 1) * size).Limit(size).Find(&orderList).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		logger.Errorf("query order list failed: %", err)
		return pageData, nil
	}
	if len(orderList) > 0 {
		copier.Copy(&resultData, &orderList)
	}
	pageData.Data = resultData
	pageData.Page = page
	pageData.PageSize = size
	pageData.Total = total
	return pageData, nil
}

// order detail info
func (o *OrderService) PayOrderDetail(id int) (node.PayOrderDetail, error) {
	var orderData order.Order
	var detailData node.PayOrderDetail
	err := o.db.Model(order.Order{}).Where("id = ?", id).First(&orderData).Error
	if err != nil {
		logger.Errorf("query pay order detail failed: %s", err)
		return detailData, err
	}
	copier.Copy(&detailData, &orderData)
	return detailData, nil
}

// cancel order
func (o *OrderService) CancelOrder(id int) error {
	var orderData order.Order
	err := o.db.Model(order.Order{}).Where("id = ?", id).First(&orderData).Error
	if err != nil {
		logger.Errorf("cancel order query order failed: %s", err)
		return err
	}
	if orderData.Status != order.PaymentPending {
		logger.Errorf("the order is not in a pending payment status and cannot be closed")
		return errors.New("the order is not in a pending payment status and cannot be closed")
	}

	err = o.db.Model(order.Order{}).Where("id = ?", id).Update("status", order.Cancelled).Error
	if err != nil {
		logger.Errorf("cancel order failed: %s", err)
		return err
	}
	return nil
}

func GenerateEthAddress() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logger.Errorf("get eth private key failed: %s", err)
		return "", "", err
	}
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	return address, privateKeyHex, nil
}

func AccountBalance(address string) (string, error) {
	client, err := ethclient.Dial(os.Getenv("NODE_URL"))
	if err != nil {
		logger.Errorf("get eth client failed: %s", err)
		return "", err
	}
	accountAddress := common.HexToAddress(address)
	erc20Contract, err := contract.NewErc20(common.HexToAddress(os.Getenv("TOKEN_ADDRESS")), client)
	if err != nil {
		logger.Errorf("get erc20 contract failed: %s", err)
		return "", err
	}
	balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		logger.Errorf("address is: %s,get balance failed: %s", address, err)
		return "", err
	}
	decimals, err := erc20Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		logger.Errorf("get token decimals failed: %s", err)
		return "", err
	}
	balanceDecimal := new(big.Float).SetInt(balance)
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	balanceDecimal.Quo(balanceDecimal, new(big.Float).SetInt(exp))
	client.Close()
	return balanceDecimal.Text('f', 2), nil
}
