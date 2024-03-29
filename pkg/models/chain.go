package models

import (
	"fmt"
	"hamster-paas/pkg/application"
	"strings"

	"gorm.io/gorm"
)

type ChainType int

const (
	Ethereum ChainType = iota
	Bsc
	Polygon
	Avalanche
	Optimism
	StarkNet
	Near
	Aptos
	Sui
	Hamster
	Ton
	Arbitrum
	Irisnet
)

func (c ChainType) String() string {
	return [...]string{"Ethereum", "Bsc", "Polygon", "Avalanche", "Optimism", "StarkNet", "Near", "Aptos", "Sui", "Hamster", "Ton", "Arbitrum", "Irisnet"}[c]
}

func (c ChainType) StringLower() string {
	return [...]string{"ethereum", "bsc", "polygon", "avalanche", "optimism", "starknet", "near", "aptos", "sui", "hamster", "ton", "arbitrum", "irisnet"}[c]
}

func ParseChainType(s string) (ChainType, error) {
	switch strings.ToLower(s) {
	case "ethereum":
		return Ethereum, nil
	case "bsc":
		return Bsc, nil
	case "polygon":
		return Polygon, nil
	case "avalanche":
		return Avalanche, nil
	case "optimism":
		return Optimism, nil
	case "starknet":
		return StarkNet, nil
	case "near":
		return Near, nil
	case "aptos":
		return Aptos, nil
	case "sui":
		return Sui, nil
	case "hamster":
		return Hamster, nil
	case "ton":
		return Ton, nil
	case "arbitrum":
		return Arbitrum, nil
	case "irisnet":
		return Irisnet, nil
	default:
		return Ethereum, fmt.Errorf("invalid chain type: %s", s)
	}
}

type RpcChain struct {
	Fullname         string `json:"fullname" gorm:"-"`
	Name             string `json:"name"`
	Network          string `json:"network"`
	HttpAddress      string `json:"httpAddress"`
	WebsocketAddress string `json:"websocketAddress"`
	ChainID          string `json:"chainID"`
	NativeToken      string `json:"nativeToken"`
	ExplorerUrl      string `json:"explorerUrl"`
	NetworkUrl       string `json:"networkUrl"`
	NetworkName      string `json:"networkName"`
	Image            string `json:"image"`
	UserActive       bool   `json:"userActive" gorm:"-"`
	IsEvm            bool   `json:"isEvm"`
	Decimals         int    `json:"decimals" gorm:"-"` // 精度，默认 18
}

type RpcChainDetail struct {
	Name   string         `json:"name"`
	Image  string         `json:"image"`
	Chains []*RpcChainApp `json:"chains"`
}

type RpcChainApp struct {
	RpcChain
	App *ApiResponseRpcApp
}

type MiddleWareRpc struct {
	RpcChain
}

func GetChainLink(chain ChainType, network NetworkType) (string, string, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return "", "", err
	}
	var c RpcChain
	err = db.Model(&RpcChain{}).Where("name = ? and network = ?", chain.StringLower(), network.StringLowerWithDash()).First(&c).Error
	if err != nil {
		return "", "", err
	}
	return c.HttpAddress, c.WebsocketAddress, nil
}

func (c ChainType) HaveNetwork(network NetworkType) bool {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return false
	}
	var count int64
	err = db.Model(&RpcChain{}).Where("name = ? and network = ?", c.StringLower(), network.StringLower()).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
