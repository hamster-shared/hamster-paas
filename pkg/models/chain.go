package models

import (
	"fmt"
	"strings"
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
	Name     string    `json:"name"`
	Networks []Network `json:"networks"`
}
