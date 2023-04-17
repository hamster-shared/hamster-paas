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
	TestnetMoonbeam
)

func (n NetworkType) String() string {
	return [...]string{"Mainnet", "TestnetGoerli", "TestnetRopsten", "TestnetKovan", "TestnetRinkeby", "TestnetSepolia", "TestnetHamster", "TestnetMumbai", "TestnetMoonbeam"}[n]
}

func (n NetworkType) StringLower() string {
	return [...]string{"mainnet", "goerli testnet", "ropsten testnet", "kovan testnet", "rinkeby testnet", "sepolia testnet", "hamster testnet", "mumbai testnet", "moonbeam testnet"}[n]
}

func (n NetworkType) StringLowerWithDash() string {
	return [...]string{"mainnet", "testnet-goerli", "testnet-ropsten", "testnet-kovan", "testnet-rinkeby", "testnet-sepolia", "testnet-hamster", "testnet-hamster", "testnet-mumbai"}[n]
}

func (n NetworkType) StringWithSpace() string {
	return [...]string{"Mainnet", "Goerli Testnet", "Ropsten Testnet", "Kovan Testnet", "Rinkeby Testnet", "Sepolia Testnet", "Hamster Testnet", "Mumbai Testnet", "Moonbeam Testnet"}[n]
}

func (n NetworkType) StringAline() string {
	return [...]string{"Mainnet", "Testnet/Goerli", "Testnet/Ropsten", "Testnet/Kovan", "Testnet/Rinkeby", "Testnet/Sepolia", "Testnet/Hamster", "Testnet/Mumbai", "Testnet/Hamster"}[n]
}

func (n NetworkType) NetworkType() eth.EthNetwork {
	return [...]eth.EthNetwork{eth.MAINNET, eth.GOERLI, eth.GOERLI, eth.GOERLI, eth.GOERLI, eth.SEPOLIA_TESTNET, eth.HAMSTER, eth.MUMBAI_TESTNET, eth.MOONBEAM_TESTNET}[n]
}

func GetNetworkIdAndUrl(network string) (string, string) {
	switch strings.ToLower(network) {
	case "sepolia testnet":
		return "aa36a7", "https://eth-sepolia.g.alchemy.com/v2/demo"
	case "mumbai testnet":
		return "13881", "https://rpc-mumbai.maticvigil.com"
	case "moonbeam testnet":
		return "501", "https://rpc-moonbeam.hamster.newtouch.com"
	default:
		return "-1", "https://rpc-moonbeam.hamster.newtouch.com"
	}
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
	case "moonbeam testnet", "testnet/moonbeam", "testnet-moonbeam":
		return TestnetMoonbeam, nil
	default:
		return Mainnet, fmt.Errorf("invalid network type: %s", s)
	}
}
