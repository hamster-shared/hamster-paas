package eth

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"hamster-paas/pkg/utils/logger"
	"testing"
	"time"
)

type EthNetwork string

const MAINNET EthNetwork = "mainnet"
const GOERLI EthNetwork = "goerli"
const HAMSTER EthNetwork = "hamster"
const BSC_MAINNET EthNetwork = "bsc_mainnet"
const BSC_TESTNET EthNetwork = "bsc_testnet"

var netMap map[EthNetwork]string = make(map[EthNetwork]string)

func setup() {
	netMap[GOERLI] = "https://goerli.infura.io/v3/ce58d7af0a4a47ec9f3d18a3545f6d18"
	netMap[MAINNET] = "https://mainnet.infura.io/v3/ce58d7af0a4a47ec9f3d18a3545f6d18"
	netMap[HAMSTER] = "https://rpc-moonbeam.hamster.newtouch.com"
	netMap[BSC_MAINNET] = "https://bsc-dataseed1.defibit.io/"
	netMap[BSC_TESTNET] = "https://data-seed-prebsc-2-s1.binance.org:8545/"
}

type networkTest struct {
	network     EthNetwork
	transaction string
}

var testCase = []networkTest{
	{MAINNET, "0x5e58cab492b6b1828d9e628390cda56a10966ab7536932d6c812f1dc882acfa4"},
	{GOERLI, "0x5b906d7dd1f0dc3644c20ad1f2369bbfa7f07638a38bf104dbd2498e76d77e8a"},
	{HAMSTER, "0x17a849c951a6f7c169872d04423d67105d174cb1564cfe20ba6af8916c636ff9"},
	{BSC_TESTNET, "0x61778dd6e01d4b6b0b693cf260b0b140a9d08c66c96be58a156bcb57c20f50d3"},
	{BSC_MAINNET, "0x161ad5781150bf35b4c4b18b5ff723dac0341852584c4766393c06d24cb321b0"},
}

func TestETHClient(t *testing.T) {
	setup()
	for _, test := range testCase {
		if _, err := GetNetworkBlockNumber(test.network); err != nil {
			t.Errorf("network %s getBlockNumber error : %v", test.network, err)
		}
	}
}

func TestGetTransaction(t *testing.T) {
	setup()
	for _, test := range testCase {
		if err := GetTransaction(test.network, test.transaction); err != nil {
			t.Errorf("network %s getBlockNumber error : %v", test.network, err)
		}
	}
}

func GetNetworkBlockNumber(network EthNetwork) (uint64, error) {

	rpcUrl := netMap[network]

	if rpcUrl == "" {
		return 0, errors.New("invalid network")
	}

	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		logger.Error("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the ", network)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	number, err := client.BlockNumber(ctx)

	return number, err
}

func GetTransaction(network EthNetwork, tx string) error {
	rpcUrl := netMap[network]

	if rpcUrl == "" {
		return errors.New("invalid network")
	}

	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		logger.Error("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the ", network)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	hash := common.Hash(common.FromHex(tx))
	transaction, isPending, err := client.TransactionByHash(ctx, hash)
	fmt.Println(transaction.Hash(), transaction.Gas())
	fmt.Println(isPending)
	return err
}
