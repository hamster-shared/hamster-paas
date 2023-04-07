package models

import (
	"fmt"
	"strings"
)

type NetworkType int

const (
	Mainnet NetworkType = iota
	TestnetGoerli
	TestnetRopsten
	TestnetKovan
	TestnetRinkeby
	TestnetSepolia
)

func (n NetworkType) String() string {
	return [...]string{"Mainnet", "TestnetGoerli", "TestnetRopsten", "TestnetKovan", "TestnetRinkeby", "TestnetSepolia"}[n]
}

func (n NetworkType) StringLower() string {
	return [...]string{"mainnet", "testnet-goerli", "testnet-ropsten", "testnet-kovan", "testnet-rinkeby", "testnet-sepolia"}[n]
}

func ParseNetworkType(s string) (NetworkType, error) {
	switch strings.ToLower(s) {
	case "mainnet":
		return Mainnet, nil
	case "testnet-goerli":
		return TestnetGoerli, nil
	case "testnet-ropsten":
		return TestnetRopsten, nil
	case "testnet-kovan":
		return TestnetKovan, nil
	case "testnet-rinkeby":
		return TestnetRinkeby, nil
	case "testnet-sepolia":
		return TestnetSepolia, nil
	default:
		return Mainnet, fmt.Errorf("invalid network type: %s", s)
	}
}
