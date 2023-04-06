package eth

import (
	"fmt"
	"testing"
)

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

func TestGetTransaction(t *testing.T) {
	setup()
	for _, test := range testCase {
		if err := GetTransaction(test.network, test.transaction); err != nil {
			t.Errorf("network %s getBlockNumber error : %v", test.network, err)
		}
	}
}

func GetTransaction(network EthNetwork, tx string) error {

	client := NewEthereumProxyFactory().GetClient(network)

	transaction, isPending, err := client.TransactionByHash(tx)
	fmt.Println(transaction.Hash(), transaction.Gas())
	fmt.Println(isPending)
	return err
}

func TestWatchRequestRes(t *testing.T) {
	client := NewEthereumProxyFactory().GetClient(HAMSTER)
	client.WatchRequestResult("0xeEA29418eBF986D5Fd18afA0005efEAC2069ac98")
}
