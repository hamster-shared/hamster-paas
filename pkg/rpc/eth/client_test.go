package eth

import (
	"fmt"
	"testing"
	"time"
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
	client := NewEthereumProxyFactory().GetClient(MUMBAI_TESTNET)
	client.WatchRequestResult("0xcb1d64Bf20faCc6093E2E39FAA7129AC8148F38a", "0x969400eb4762e926adf55aa3deebfbd6b32db30e8dc866fa3fe0ca7db47827d3", "")
}

func TestGetReceipt(t *testing.T) {
	client := NewEthereumProxyFactory().GetClient(MUMBAI_TESTNET)
	receipt, err := client.TransactionReceipt("0x71db36d509f640b8cf103f8e695cf1621f0d2b5dadb51efc5a3ec9d649088aba")
	if err != nil {
		panic(err)
	}
	fmt.Println(receipt.Status)
}

// 直接拿的时间比重新链接client的时间要短
func TestGetChainClient(t *testing.T) {
	begin := time.Now()
	client := GetChainClient(MUMBAI_TESTNET)
	if client == nil {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
	end := time.Now()
	begin2 := time.Now()
	client2 := GetChainClient(MUMBAI_TESTNET)
	if client2 == nil {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
	end2 := time.Now()
	fmt.Println("1 : ", end.Second()-begin.Second())
	fmt.Println("2 : ", end2.Second()-begin2.Second())
}

func TestGetTxStatus(t *testing.T) {
	client := GetChainClient(SEPOLIA_TESTNET)
	if client == nil {
		panic("client not valid")
	}
	status, err := GetTxStatus("0xd02009d5f2a521ee2ce9bec2de8c79b00ee59e426e6e17c4012c455e91ea118b", MOONBEAM_TESTNET, client)
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
}
