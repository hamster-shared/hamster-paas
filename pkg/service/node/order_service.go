package node

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	modelsNode "hamster-paas/pkg/models/node"
	"hamster-paas/pkg/models/order"
	"hamster-paas/pkg/models/vo/node"
	"hamster-paas/pkg/utils/logger"
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

func (o *OrderService) LaunchOrder(userId int, launchData node.LaunchOrderParam) (uint, error) {
	var changeAccount modelsNode.UserChargeAccount
	var orderData order.Order
	err := o.db.Model(modelsNode.UserChargeAccount{}).Where("user_id = ?", userId).First(&changeAccount).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			address, privateKey, err := GenerateEthAddress()
			if err != nil {
				logger.Errorf("generate address failed: %s", err)
				return 0, err
			}
			changeAccount.UserId = userId
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
	orderId := o.snowFlakeNode.Generate().String()
	orderData.OrderId = orderId
	orderData.OrderTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	orderData.OrderType = order.NodeService
	orderData.UserId = uint(userId)
	orderData.Status = order.PaymentPending
	orderData.Chain = modelsNode.ChainProtocol(launchData.Protocol)
	orderData.Amount = launchData.Amount
	orderData.BuyTime = launchData.BuyTime
	orderData.ResourceType = launchData.ResourceType
	orderData.ReceiveAddress = changeAccount.Address
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
func (o *OrderService) OrderList(startDateStr, endDateStr, query string, userId, page, size int) (node.OrderPage, error) {
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
	err := tx.Offset((page - 1) * size).Limit(size).Find(&orderList).Offset(-1).Limit(-1).Count(&total).Error
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
