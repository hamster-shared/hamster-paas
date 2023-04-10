package eth

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"hamster-paas/pkg/consts"
	"hamster-paas/pkg/utils/logger"
	"math/big"
	"strconv"
	"strings"
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
	//netMap[HAMSTER] = "https://rpc-moonbeam.hamster.newtouch.com"
	//netMap[HAMSTER] = "wss://ws-moonbeam.hamster.newtouch.com"
	//netMap[HAMSTER] = "wss://eth-sepolia.g.alchemy.com/v2/BgE-iyk7FqwXwyn6pHEeByyZpI56NYgO"
	netMap[HAMSTER] = "wss://polygon-mumbai.g.alchemy.com/v2/BM4kwUJwMKmdh1zaDDByzNr19jgzdRiV"
	netMap[BSC_MAINNET] = "https://bsc-dataseed1.defibit.io/"
	netMap[BSC_TESTNET] = "https://data-seed-prebsc-2-s1.binance.org:8545/"
}

func init() {
	setup()
}

type EthereumProxy interface {
	TransactionByHash(hash string) (tx *types.Transaction, isPending bool, err error)
	WatchRequestResult(contractAddress, requestId string) error
	TransactionReceipt(hash string) (*types.Receipt, error)
}

type EthereumProxyFactory struct {
	clients map[EthNetwork]EthereumProxy
}

func NewEthereumProxyFactory() *EthereumProxyFactory {

	return &EthereumProxyFactory{
		clients: make(map[EthNetwork]EthereumProxy),
	}
}

// GetClient  get ethereum client
func (e *EthereumProxyFactory) GetClient(network EthNetwork) EthereumProxy {
	if val, ok := e.clients[network]; ok {
		//do
		return val
	} else {
		client, err := NewRPCEthereumProxy(netMap[network])
		if err != nil {
			return nil
		}
		e.clients[network] = client
		return client
	}
}

type RPCEthereumProxy struct {
	client *ethclient.Client
}

func NewRPCEthereumProxy(endpoints string) (*RPCEthereumProxy, error) {
	client, err := ethclient.Dial(endpoints)

	if err != nil {
		logger.Error("Oops! There was a problem", err)
		return nil, err
	} else {
		fmt.Println("Success! you are connected to the ", endpoints)
		return &RPCEthereumProxy{
			client,
		}, nil
	}
}

func (rpc *RPCEthereumProxy) TransactionByHash(hash string) (tx *types.Transaction, isPending bool, err error) {
	ctx := context.Background()
	hashTx := common.Hash(common.FromHex(hash))
	return rpc.client.TransactionByHash(ctx, hashTx)
}

type OCRResponseEvent struct {
	RequestId [32]byte
	Result    []byte
	Err       []byte
}

func (rpc *RPCEthereumProxy) WatchRequestResult(contractAddress, requestId string) error {
	// 要监听的合约地址
	oracleContractAddress := common.HexToAddress(contractAddress)
	//consumerContract, err := contract2.NewFunctionConsumer(oracleContractAddress, rpc.client)
	//if err != nil {
	//	fmt.Println("get contract failed:", err.Error())
	//	return err
	//}
	//var requestIdBytes [32]byte
	//copy((*[32]byte)(unsafe.Pointer(&requestIdBytes[0]))[:], requestId)
	//var array [][32]byte
	//array = append(array, requestIdBytes)
	//eventChan := make(chan *contract2.FunctionConsumerOCRResponse)
	//opts := &bind.WatchOpts{Context: context.Background()}
	//sub, err := consumerContract.WatchOCRResponse(opts, eventChan, array)
	//for {
	//	select {
	//	case event := <-eventChan:
	//		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//		fmt.Printf("sender: %s, value: %s\n", event.Result, event.RequestId)
	//		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//	case err := <-sub.Err():
	//		fmt.Println(err.Error())
	//	}
	//}
	contractAbi, err := abi.JSON(strings.NewReader(consts.ConsumerAbi))
	// 监听请求结果
	query := ethereum.FilterQuery{
		Addresses: []common.Address{oracleContractAddress},
		Topics: [][]common.Hash{
			{
				contractAbi.Events["OCRResponse"].ID,
				//crypto.Keccak256Hash([]byte("OCRResponse(bytes32,bytes,bytes)")),
				// 用于监听请求结果的事件签名
				//common.HexToHash(consts.EventSignature),
			},
		},
	}
	logs := make(chan types.Log)
	sub, err := rpc.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		//logger.Error("Error subscribing to logs:", err)
		fmt.Println("--------------------------")
		fmt.Println(err.Error())
		fmt.Println("--------------------------")
		return err
	}

	// 处理请求结果
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Error in subscription:", err)
			return errors.New(fmt.Sprintf("Error in subscription:%s", err.Error()))
		case log := <-logs:
			fmt.Println("++++++++++++++++++++++++++++++++++++++")
			fmt.Println(log.Address)
			fmt.Println(log.Data)
			fmt.Println(log.Topics)
			fmt.Println(log.BlockHash)
			fmt.Println(log.Index)
			fmt.Println(log.BlockNumber)
			fmt.Println(log.TxIndex)
			fmt.Println(log.TxHash)
			if err != nil {
				fmt.Println("---------------------------------------")
				fmt.Println("Failed to parse ABI:", err.Error())
				fmt.Println("---------------------------------------")
			}

			// 解析事件数据
			var eventData OCRResponseEvent
			err = contractAbi.UnpackIntoInterface(&eventData, "OCRResponse", log.Data)
			if err != nil {
				fmt.Println("---------------------------------------")
				fmt.Println("Failed to unpack event data: ", err.Error())
				fmt.Println("---------------------------------------")
			}
			fmt.Println("++++++++++++++++++++++++++++++++++++++")
			fmt.Println("****************************************")
			fmt.Printf("RequestId: %x\n", eventData.RequestId)
			fmt.Printf("Result: %s\n", hexToString(eventData.Result))
			data, _ := strconv.Atoi(string(eventData.Result))
			fmt.Println(data)
			fmt.Printf("Err: %s\n", string(eventData.Err))
			fmt.Println("****************************************")
			// 比对 Request ID
			requestID := log.Topics[1].Big()
			expectedRequestID := big.NewInt(123)
			if requestID.Cmp(expectedRequestID) == 0 {
				fmt.Println("data")
			}
		}
	}
	return nil
}

func hexToString(hex []byte) string {
	s := ""
	for _, b := range hex {
		s += fmt.Sprintf("%02x", b)
	}
	return s
}

// TransactionReceipt 通过Receipt.status来判断交易事务状态，0 -> 失败， 1 -> 成功
func (rpc *RPCEthereumProxy) TransactionReceipt(hash string) (*types.Receipt, error) {
	hashTx := common.Hash(common.FromHex(hash))
	return rpc.client.TransactionReceipt(context.Background(), hashTx)
}
