package models

import (
	"fmt"
	"hamster-paas/pkg/application"
	"strings"

	"gorm.io/gorm"
)

type ChainType int

const (
	Ethereum ChainType = iota + 1
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
	return [...]string{"Ethereum", "Bsc", "Polygon", "Avalanche", "Optimism", "StarkNet", "Near", "Aptos", "Sui"}[c-1]
}

func (c ChainType) StringLower() string {
	return [...]string{"ethereum", "bsc", "polygon", "avalanche", "optimism", "starknet", "near", "aptos", "sui"}[c-1]
}

func ParseChainType(s string) (ChainType, error) {
	switch strings.ToLower(s) {
	case "ethereum", "Ethereum":
		return Ethereum, nil
	case "bsc", "Bsc":
		return Bsc, nil
	case "polygon", "Polygon":
		return Polygon, nil
	case "avalanche", "Avalanche":
		return Avalanche, nil
	case "optimism", "Optimism":
		return Optimism, nil
	case "starknet", "Starknet":
		return StarkNet, nil
	case "near", "Near":
		return Near, nil
	case "aptos", "Aptos":
		return Aptos, nil
	case "sui", "Sui":
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

type RpcChainDetail struct {
	RpcChainBaseInfo
	Chains []*RpcChainApp `json:"chains"`
}

type RpcChainBaseInfo struct {
	ChainID     int    `json:"chain_id"`
	NativeToken string `json:"native_token"`
	Explorers   string `json:"explorers"`
}

type RpcChainApp struct {
	RpcChain
	App *ApiResponseRpcApp
}

type MiddleWareRpc struct {
	RpcChainBaseInfo
	RpcChain
}

func GetChainLink(chain ChainType, network NetworkType) (string, string, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return "", "", err
	}
	var c RpcChain
	err = db.Model(&RpcChain{}).Where("name = ? and network = ?", chain.StringLower(), network.StringLower()).First(&c).Error
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

func (c ChainType) BaseInfo() RpcChainBaseInfo {
	switch c {
	case Ethereum:
		return getEthereumBaseInfo()
	case Bsc:
		return getBSCBaseInfo()
	case Polygon:
		return getPolygonBaseInfo()
	case Avalanche:
		return getAvalancheBaseInfo()
	case Optimism:
		return getOptimismBaseInfo()
	case StarkNet:
		return getStarkNetBaseInfo()
	case Near:
		return getNEARBaseInfo()
	case Aptos:
		return getAptosBaseInfo()
	case Sui:
		return getSUIBaseInfo()
	default:
		return getEthereumBaseInfo()
	}
}

func getEthereumBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Ethereum),
		NativeToken: "ETH",
		Explorers:   "https://etherscan.io/",
	}
}

func getBSCBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Bsc),
		NativeToken: "BNB",
		Explorers:   "https://bscscan.com/",
	}
}

func getPolygonBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Polygon),
		NativeToken: "MATIC",
		Explorers:   "https://polygonscan.com/",
	}
}

func getAvalancheBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Avalanche),
		NativeToken: "AVAX",
		Explorers:   "https://cchain.explorer.avax.network/",
	}
}

func getOptimismBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Optimism),
		NativeToken: "ETH",
		Explorers:   "https://optimistic.etherscan.io/",
	}
}

func getStarkNetBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(StarkNet),
		NativeToken: "STARK",
		Explorers:   "https://voyager.online/",
	}
}

func getNEARBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Near),
		NativeToken: "NEAR",
		Explorers:   "https://explorer.near.org/",
	}
}

func getAptosBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Aptos),
		NativeToken: "APTO",
		Explorers:   "https://explorer.aptos.network/",
	}
}

func getSUIBaseInfo() RpcChainBaseInfo {
	return RpcChainBaseInfo{
		ChainID:     int(Sui),
		NativeToken: "SUI",
		Explorers:   "https://explorer.sovryn.app/",
	}
}
