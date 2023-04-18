package models

import (
	"fmt"
	"strings"
)

func GetAlineNetwork(chain string, network string) (string, error) {
	chainNetwork := chain + " " + network
	switch strings.ToLower(chainNetwork) {
	case "ethereum mainnet":
		return "Ethereum/Mainnet", nil
	case "ethereum goerli testnet":
		return "Ethereum/Goerli", nil
	case "ethereum sepolia testnet":
		return "Ethereum/Sepolia", nil
	case "ethereum hamster testnet":
		return "Ethereum/Hamster", nil
	case "polygon mainnet":
		return "Polygon/Mainnet", nil
	case "polygon mumbai testnet":
		return "Polygon/Mumbai", nil
	case "bsc mainnet":
		return "Bsc/Mainnet", nil
	case "bsc testnet":
		return "Bsc/Testnet", nil
	case "hamster moonbeam testnet":
		return "Ethereum/Hamster", nil
	default:
		return "", fmt.Errorf("chain and network invalid: %v, %v", chain, network)
	}
}

func AlineNetworkToChainkLinkNetwork(alineNetwork string) string {
	switch alineNetwork {
	case "Ethereum/Mainnet":
		return "Ethereum Mainnet"
	case "Ethereum/Goerli":
		return "Ethereum Goerli Testnet"
	case "Ethereum/Sepolia":
		return "Ethereum Sepolia Testnet"
	case "Ethereum/Hamster":
		return "Hamster Moonbeam Testnet"
	case "Polygon/Mainnet":
		return "Polygon Mainnet"
	case "Polygon/Mumbai":
		return "Polygon Mumbai Testnet"
	case "Bsc/Mainnet":
		return "Bsc Mainnet"
	case "Bsc/Testnet":
		return "Bsc Testnet"
	default:
		return "Network Unknown"
	}
}

func GetChainFrameType(alineNetwork string) string {
	switch alineNetwork {
	case "Ethereum/Mainnet":
		return "1"
	case "Ethereum/Goerli":
		return "1"
	case "Ethereum/Sepolia":
		return "1"
	case "Ethereum/Hamster":
		return "1"
	case "Polygon/Mainnet":
		return "1"
	case "Polygon/Mumbai":
		return "1"
	case "Bsc/Mainnet":
		return "1"
	case "Bsc/Testnet":
		return "1"
	default:
		return "Network Unknown"
	}
}
