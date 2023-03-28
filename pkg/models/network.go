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
)

func (n NetworkType) String() string {
	return [...]string{"Mainnet", "TestnetGoerli", "TestnetRopsten", "TestnetKovan", "TestnetRinkeby"}[n]
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
	default:
		return Mainnet, fmt.Errorf("invalid network type: %s", s)
	}
}
