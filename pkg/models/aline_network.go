package models

import (
	"strings"
)

func GetAlineNetwork(chain string, network string) string {
	chainNetwork := chain + " " + network
	switch strings.ToLower(chainNetwork) {
	case "ethereum mainnet":
		return "Ethereum/Mainnet"
	case "ethereum goerli testnet":
		return "Ethereum/Goerli"
	case "ethereum sepolia testnet":
		return "Ethereum/Sepolia"
	case "ethereum hamster testnet":
		return "Ethereum/Hamster"
	case "polygon mainnet":
		return "Polygon/Mainnet"
	case "polygon mumbai testnet":
		return "Polygon/Mumbai"
	case "bsc mainnet":
		return "Bsc/Mainnet"
	case "bsc testnet":
		return "Bsc/Testnet"
	case "hamster moonbeam testnet":
		return "Ethereum/Hamster"
	default:
		return network
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
		return alineNetwork
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
