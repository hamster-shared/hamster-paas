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

func GetChains() ([]Chain, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var chains []Chain
	err = db.Model(&Chain{}).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	return chains, nil
}

func GetNetworks(chain ChainType) ([]string, error) {
	db, err := application.GetBean[*gorm.DB]("db")
	if err != nil {
		return nil, err
	}
	var chains []Chain
	err = db.Model(&Chain{}).Where("name = ?", chain.String()).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	var networks []string
	for _, chain := range chains {
		networks = append(networks, chain.Network)
	}
	return networks, nil
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

type Chain struct {
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
	var c Chain
	err = db.Model(&Chain{}).Where("name = ? and network = ?", chain.String(), network.String()).First(&c).Error
	if err != nil {
		return "", "", err
	}
	return c.HttpAddress, c.WebsocketAddress, nil
}
