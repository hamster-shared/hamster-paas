package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"hamster-paas/pkg/service/contract"
	"hamster-paas/pkg/utils"
	"hamster-paas/pkg/utils/logger"
	"log"
	"strconv"
	"time"
)

type EthNetwork string

const MAINNET EthNetwork = "mainnet"
const GOERLI EthNetwork = "Goerli Testnet"
const HAMSTER EthNetwork = "hamster"
const BSC_MAINNET EthNetwork = "bsc_mainnet"
const BSC_TESTNET EthNetwork = "bsc_testnet"
const SEPOLIA_TESTNET EthNetwork = "Sepolia Testnet"
const MUMBAI_TESTNET EthNetwork = "Mumbai Testnet"
const RINKBEY_TESTNET EthNetwork = "Rinkeby Testnet"
const MOONBEAM_TESTNET EthNetwork = "Moonbeam Testnet"

var NetMap map[EthNetwork]string = make(map[EthNetwork]string)
var ClientMap map[EthNetwork]*ethclient.Client = make(map[EthNetwork]*ethclient.Client)

func setup() {
	NetMap[GOERLI] = "https://goerli.infura.io/v3/ce58d7af0a4a47ec9f3d18a3545f6d18"
	NetMap[MAINNET] = "https://mainnet.infura.io/v3/ce58d7af0a4a47ec9f3d18a3545f6d18"
	NetMap[HAMSTER] = "wss://polygon-mumbai.g.alchemy.com/v2/BM4kwUJwMKmdh1zaDDByzNr19jgzdRiV"
	NetMap[BSC_MAINNET] = "https://bsc-dataseed1.defibit.io/"
	NetMap[BSC_TESTNET] = "https://data-seed-prebsc-2-s1.binance.org:8545/"
	NetMap[MUMBAI_TESTNET] = "wss://polygon-mumbai.g.alchemy.com/v2/DR5idsZrnqA4c3gegtjOjlVHLkGkLLMn"
	NetMap[SEPOLIA_TESTNET] = "wss://eth-sepolia.g.alchemy.com/v2/WqkvURejyuhwy2WZagmne0VqRdunMpEd"
	NetMap[MOONBEAM_TESTNET] = "wss://ws-moonbeam.hamster.newtouch.com"
}

func init() {
	setup()
}

type EthereumProxy interface {
	TransactionByHash(hash string) (tx *types.Transaction, isPending bool, err error)
	WatchRequestResult(contractAddress, requestId, email, requestName string) error
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
		client, err := NewRPCEthereumProxy(NetMap[network])
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

func (rpc *RPCEthereumProxy) WatchRequestResult(contractAddress, requestId, email, requestName string) error {
	log.Println("++++++++++++++++++++++++++++++++++++++")
	// 要监听的合约地址
	oracleContractAddress := common.HexToAddress(contractAddress)
	// 定义查询过滤器
	query := ethereum.FilterQuery{
		Addresses: []common.Address{oracleContractAddress},
		Topics: [][]common.Hash{
			{
				crypto.Keccak256Hash([]byte("OCRResponse(bytes32,bytes,bytes)")),
			},
		},
	}
	// 创建订阅
	logs := make(chan types.Log)
	sub, err := rpc.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Errorf("create Mumbai subscribe Oracle Request failed: %s", err)
	}
	contractFilter, err := contract.NewFunctionConsumer(oracleContractAddress, rpc.client)
	if err != nil {
		logger.Errorf("create Oracle Request failed: %s", err)
	}
	// 监听订阅事件
	for {
		select {
		case err := <-sub.Err():
			logger.Errorf("subscribe Oracle Request event failed: %s", err)
		case vLog := <-logs:
			logger.Info("start watch Oracle Request event send email")
			data, err := contractFilter.ParseOCRResponse(vLog)
			if err == nil {
				requestIdData := fmt.Sprintf("0x%s", hex.EncodeToString(data.RequestId[:]))
				log.Println("++++++++++++++++++++")
				fmt.Printf("request id is:%s", requestIdData)
				log.Println("++++++++++++++++++++")
				if requestIdData == requestId {
					var result string
					numData, err := strconv.ParseInt(hexToString(data.Result), 16, 64)
					if err != nil {
						result = string(data.Result)
					} else {
						result = strconv.Itoa(int(numData))
					}
					utils.SendEmail(email, requestId, result, requestName, string(data.Err))
					break
				}
			} else {
				logger.Errorf("parse OracleRequest data failed: %s", err)
			}
		}
	}

	//contractAbi, err := abi.JSON(strings.NewReader(consts.ConsumerAbi))
	//// 监听请求结果
	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{oracleContractAddress},
	//	Topics: [][]common.Hash{
	//		{
	//			contractAbi.Events["OCRResponse"].ID,
	//		},
	//	},
	//}
	//logs := make(chan types.Log)
	//sub, err := rpc.client.SubscribeFilterLogs(context.Background(), query, logs)
	//if err != nil {
	//	//logger.Error("Error subscribing to logs:", err)
	//	fmt.Printf("query logs failed: %s\n", err.Error())
	//	return err
	//}
	//
	//// 处理请求结果
	//for {
	//	select {
	//	case err := <-sub.Err():
	//		fmt.Println("Error in subscription:", err)
	//		return errors.New(fmt.Sprintf("Error in subscription:%s", err.Error()))
	//	case log := <-logs:
	//		// 解析事件数据
	//		var eventData OCRResponseEvent
	//		err = contractAbi.UnpackIntoInterface(&eventData, "OCRResponse", log.Data)
	//		if err != nil {
	//			fmt.Println("Failed to unpack event data: ", err.Error())
	//			return err
	//		}
	//		if len(log.Topics) == 2 {
	//			fmt.Printf("RequestId: %x\n", log.Topics[1])
	//			if requestId == log.Topics[1].String() {
	//				var result string
	//				numData, err := strconv.ParseInt(hexToString(eventData.Result), 16, 64)
	//				if err != nil {
	//					result = string(eventData.Result)
	//				} else {
	//					result = strconv.Itoa(int(numData))
	//				}
	//				utils.SendEmail(email, requestId, result, string(eventData.Err))
	//				break
	//			}
	//		}
	//	}
	//}
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

// GetChainClient 获取chain的client
func GetChainClient(ethNetwork EthNetwork) *ethclient.Client {
	var client *ethclient.Client
	var err error
	var ok bool
	if client, ok = ClientMap[ethNetwork]; ok {
		if client != nil {
			_, err = client.NetworkID(context.Background())
			if err == nil {
				return client
			}
		}
	}
	times := 0
	for {
		if times == 10 {
			break
		}
		client, err = ethclient.Dial(NetMap[ethNetwork])
		// 连接成功，插入到ClientMap中
		if err == nil {
			ClientMap[ethNetwork] = client
			//logger.Infof("chain client：%v 重新连接或连接失效，重新链接成功", ethNetwork)
			return client
		}
		time.Sleep(time.Second * 5)
		times++
	}
	return nil
}

// GetTxStatus 获取交易状态
func GetTxStatus(hash string, ethNetwork EthNetwork, client *ethclient.Client) (*types.Receipt, error) {
	if len(hash) != 66 {
		return nil, fmt.Errorf("hash length error")
	}
	r, err := client.TransactionReceipt(context.Background(), common.Hash(common.FromHex(hash)))
	if err != nil {
		return nil, fmt.Errorf("get tx receipt faild")
	}
	return r, nil
}
