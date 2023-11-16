package service

import (
	"github.com/jinzhu/copier"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"hamster-paas/pkg/service/zan"

	"gorm.io/gorm"
)

type MiddleWareService struct {
	zanClient *zan.ZanClient
	db        *gorm.DB
}

func NewMiddleWareService(db *gorm.DB, zanClient *zan.ZanClient) *MiddleWareService {
	return &MiddleWareService{
		db:        db,
		zanClient: zanClient,
	}
}

func (s *MiddleWareService) MiddleWareRpc(userID string) ([]vo.MiddleWareRpcZan, error) {
	var middleWareRpcZanList []vo.MiddleWareRpcZan
	buyFlag := false
	digest, err := s.zanClient.EcosystemsDigest()
	if err != nil {
		return middleWareRpcZanList, err
	}
	var zanUser models.ZanUser
	err = s.db.Model(&models.ZanUser{}).Where("user_id = ?", userID).First(&zanUser).Error
	if err == nil {
		apiKeyList, err := s.zanClient.ApiKeyList(1, 10, zanUser.AccessToken)
		if err == nil && len(apiKeyList.Data.Data) > 0 {
			buyFlag = true
		}
	}

	nameTokenSymbol := make(map[string]string)
	nameTokenSymbol["Ethereum"] = "ETH"
	nameTokenSymbol["BSC"] = "BNB"
	nameTokenSymbol["Polygon"] = "MATIC"
	nameTokenSymbol["Arbitrum"] = "ARB"
	nameTokenSymbol["Optimism"] = "OP"
	nameTokenSymbol["Tron"] = "TRX"
	nameTokenSymbol["Sui"] = "SUI"
	nameTokenSymbol["Aptos"] = "APT"
	nameTokenSymbol["Bitcoin"] = "BTC"
	nameTokenSymbol["Avalanche"] = "AVAX"
	for _, digestInfo := range digest.Data {
		var middleWareRpcZan vo.MiddleWareRpcZan
		copier.Copy(&middleWareRpcZan, &digestInfo)
		middleWareRpcZan.BuyFlag = buyFlag
		if symbol, ok := nameTokenSymbol[digestInfo.EcosystemName]; ok {
			middleWareRpcZan.EcosystemCode = symbol
		}
		middleWareRpcZanList = append(middleWareRpcZanList, middleWareRpcZan)
	}
	return middleWareRpcZanList, nil
}
