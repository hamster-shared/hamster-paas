package models

import (
	"fmt"
	"hamster-paas/pkg/rpc/eth"
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
	TestnetHamster
	TestnetMumbai
)

func (n NetworkType) String() string {
	return [...]string{"Mainnet", "TestnetGoerli", "TestnetRopsten", "TestnetKovan", "TestnetRinkeby", "TestnetSepolia", "TestnetHamster", "TestnetMumbai"}[n]
}

func (n NetworkType) StringLower() string {
	return [...]string{"mainnet", "goerli testnet", "ropsten testnet", "kovan testnet", "rinkeby testnet", "sepolia testnet", "hamster testnet", "mumbai testnet"}[n]
}

func (n NetworkType) StringWithSpace() string {
	return [...]string{"Mainnet", "Goerli Testnet", "Ropsten Testnet", "Kovan Testnet", "Rinkeby Testnet", "Sepolia Testnet", "Hamster Testnet", "Mumbai Testnet"}[n]
}

func (n NetworkType) StringAline() string {
	return [...]string{"Mainnet", "Testnet/Goerli", "Testnet/Ropsten", "Testnet/Kovan", "Testnet/Rinkeby", "Testnet/Sepolia", "Testnet/Hamster", "Mumbai Testnet"}[n]
}

func (n NetworkType) NetworkType() eth.EthNetwork {
	return [...]eth.EthNetwork{eth.MAINNET, eth.GOERLI, eth.GOERLI, eth.GOERLI, eth.GOERLI, eth.GOERLI, eth.GOERLI, eth.MUMBAI_TESTNET}[n]
}

func ParseNetworkType(s string) (NetworkType, error) {
	switch strings.ToLower(s) {
	case "mainnet":
		return Mainnet, nil
	case "goerli testnet", "testnet/goerli", "testnet-goerli":
		return TestnetGoerli, nil
	case "ropsten testnet", "testnet/ropsten", "testnet-ropsten":
		return TestnetRopsten, nil
	case "kovan testnet", "testnet/kovan", "testnet-kovan":
		return TestnetKovan, nil
	case "rinkeby testnet", "testnet/rinkeby", "testnet-rinkeby":
		return TestnetRinkeby, nil
	case "sepolia testnet", "testnet/sepolia", "testnet-sepolia":
		return TestnetSepolia, nil
	case "hamster testnet", "testnet/hamster", "testnet-hamster":
		return TestnetHamster, nil
	case "mumbai testnet", "testnet/mumbai", "testnet-mumbai":
		return TestnetMumbai, nil
	default:
		return Mainnet, fmt.Errorf("invalid network type: %s", s)
	}
}
