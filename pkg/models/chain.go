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
)

func (c ChainType) String() string {
	return [...]string{"Ethereum", "Bsc", "Polygon", "Avalanche", "Optimism", "StarkNet", "Near", "Aptos", "Sui"}[c]
}

func (c ChainType) StringLower() string {
	return [...]string{"ethereum", "bsc", "polygon", "avalanche", "optimism", "starknet", "near", "aptos", "sui"}[c]
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
	default:
		return Ethereum, fmt.Errorf("invalid chain type: %s", s)
	}
}

type RpcChain struct {
	Name             string `json:"name"`
	Network          string `json:"network"`
	HttpAddress      string `json:"http_address"`
	WebsocketAddress string `json:"websocket_address"`
}

func GetChainLink(chain ChainType, network NetworkType) (string, string, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return "", "", err
	}
	var c RpcChain
	err = db.Model(&RpcChain{}).Where("name = ? and network = ?", chain.String(), network.String()).First(&c).Error
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
