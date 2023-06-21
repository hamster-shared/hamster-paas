package node

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	modelsNode "hamster-paas/pkg/models/node"
	"hamster-paas/pkg/models/vo/node"
	"hamster-paas/pkg/utils/logger"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NodeService struct {
	db *gorm.DB
}

func NewNodeService(db *gorm.DB) *NodeService {
	return &NodeService{
		db: db,
	}
}

// node statistic
func (n *NodeService) NodeStatisticsInfo(userId int) (node.NodeStatisticsVo, error) {
	var statisticsInfo node.NodeStatisticsVo
	results := []struct {
		Status int
		Count  int
	}{}
	err := n.db.Model(modelsNode.RPCNode{}).Select("status, count(*) as count").Where("user_id = ?", userId).
		Group("status").
		Scan(&results).
		Error
	if err != nil {
		logger.Errorf("query node statistics info failed: %s", err)
		return statisticsInfo, err
	}
	if len(results) > 0 {
		for _, result := range results {
			if result.Status == int(modelsNode.Synced) {
				statisticsInfo.Synced = result.Count
			} else {
				statisticsInfo.Halted = result.Count
			}
			statisticsInfo.Nodes = statisticsInfo.Nodes + int64(result.Count)
		}
	}
	return statisticsInfo, nil
}

// node list
func (n *NodeService) NodeList(userId, page, size int) (node.NodePage, error) {
	var nodes []modelsNode.RPCNode
	var data []node.NodeVo
	var total int64
	var nodePage node.NodePage
	err := n.db.Model(modelsNode.RPCNode{}).Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&nodes).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		logger.Errorf("query node  list failed: %s", err)
		return nodePage, err
	}
	if len(nodes) > 0 {
		for _, nodeData := range nodes {
			var nodeVo node.NodeVo
			copier.Copy(&nodeVo, &nodeData)
			if nodeData.HttpEndpoint == "" {
				data = append(data, nodeVo)
				continue
			}
			status, _ := QueryNodeStatus(nodeData.HttpEndpoint, string(nodeData.ChainProtocol))
			nodeData.Status = status
			data = append(data, nodeVo)
		}
	}
	nodePage.Data = data
	nodePage.Page = page
	nodePage.PageSize = size
	nodePage.Total = total
	return nodePage, nil
}

// node detail
func (n *NodeService) NodeDetail(nodeId int) (node.NodeDetail, error) {
	var nodeData modelsNode.RPCNode
	var detailData node.NodeDetail
	err := n.db.Model(modelsNode.RPCNode{}).Where("id = ?", nodeId).First(&nodeData).Error
	if err != nil {
		logger.Errorf("query node  detail failed: %s", err)
		return detailData, err
	}
	copier.Copy(&detailData, &nodeData)
	if detailData.HttpEndpoint != "" {
		detailData = QueryNodeInfo(detailData)
		switch string(detailData.ChainProtocol) {
		case "Ethereum", "Optimism", "Avalanche":
			blockTime := QueryBlockTime(detailData)
			detailData.BlockTime = blockTime
		}
	}
	return detailData, nil
}

func (n *NodeService) UpdateNode(id uint, nodeData node.UpdateNodeParam) error {
	var node modelsNode.RPCNode
	err := n.db.Where("id? and user_id=?", id).First(&node).Error
	if err != nil {
		return fmt.Errorf("node id: %d is aready exists", id)
	}
	node.PublicIp = nodeData.PublicIp
	node.ChainVersion = string(nodeData.ChainVersion)
	node.Status = nodeData.Status
	node.RemainingSyncTime = nodeData.RemainingSyncTime
	node.HttpEndpoint = nodeData.HttpEndpoint
	node.WebsocketEndpoint = nodeData.WebsocketEndpoint
	err = n.db.Model(&node).Updates(&node).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryNodeStatus(rpcUrl, chainProtocol string) (modelsNode.RPCNodeStatus, error) {
	switch chainProtocol {
	case "Ethereum", "Optimism":
		client, err := rpc.Dial(rpcUrl)
		if err != nil {
			logger.Errorf("query node status get client failed: %s", err)
			return modelsNode.Halted, err
		}
		defer client.Close()
		var state interface{}
		err = client.CallContext(context.Background(), &state, "eth_syncing")
		if err != nil {
			logger.Errorf("get node sync status failed: %s", err)
			return modelsNode.Halted, err
		}
		if state == "false" {
			return modelsNode.Synced, nil
		} else {
			return modelsNode.Syncing, nil
		}
	case "Starknet":
		payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"starknet_syncing\"}")
		req, _ := http.NewRequest("POST", rpcUrl, payload)
		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		result := struct {
			Id      int64       `json:"id"`
			JsonRpc string      `json:"jsonrpc"`
			Result  interface{} `json:"result"`
		}{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			logger.Errorf("get stark net status failed: %s", err)
			return modelsNode.Halted, err
		}
		if b, ok := result.Result.(bool); ok {
			if b == false {
				return modelsNode.Synced, nil
			}
		} else {
			return modelsNode.Syncing, nil
		}
	case "Near":
		versionPayload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"status\"}")
		req, _ := http.NewRequest("POST", rpcUrl, versionPayload)
		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		result := struct {
			Id      int64            `json:"id"`
			JsonRpc string           `json:"jsonrpc"`
			Result  NearStatusResult `json:"result"`
		}{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			logger.Errorf("near status json unmarshal failed: %s", err)
			return modelsNode.Halted, err
		}
		if result.Result.SyncInfo.Syncing == false {
			return modelsNode.Synced, nil
		}
		return modelsNode.Syncing, nil
	}
	return modelsNode.Initializing, nil
}

type NearStatusResult struct {
	SyncInfo NearSyncInfo `json:"sync_info"`
}
type NearSyncInfo struct {
	Syncing bool `json:"syncing"`
}

func QueryNodeInfo(data node.NodeDetail) node.NodeDetail {
	status, _ := QueryNodeStatus(data.HttpEndpoint, string(data.ChainProtocol))
	data.Status = status
	switch string(data.ChainProtocol) {
	case "Ethereum", "Optimism", "Avalanche":
		client, err := rpc.Dial("https://eth-mainnet.g.alchemy.com/v2/WF6iALbJ7XE4XR3MWvgUvIo2NtkuJzgK")
		if err != nil {
			logger.Errorf("query node info failed: %s", err)
			return data
		}
		defer client.Close()
		var blockNumber string
		err = client.CallContext(context.Background(), &blockNumber, "eth_blockNumber")
		if err == nil {
			dec := new(big.Int)
			dec.SetString(blockNumber[2:], 16)
			data.CurrentHeight = uint(dec.Uint64())
		} else {
			logger.Errorf("query node block number failed: %s", err)
		}
		//query version
		var version string
		err = client.Call(&version, "web3_clientVersion")
		if err != nil {
			logger.Errorf("query node version failed: %s", err)
			return data
		}
		data.ChainVersion = version
	case "Starknet":
		payload := strings.NewReader("{\"id\":1,\"jsonrpc\":\"2.0\",\"method\":\"starknet_blockNumber\"}")

		req, _ := http.NewRequest("POST", data.HttpEndpoint, payload)

		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		result := struct {
			Id      int64  `json:"id"`
			JsonRpc string `json:"jsonrpc"`
			Result  int64  `json:"result"`
		}{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			logger.Errorf("stark net block json unmarshal failed: %s", err)
			return data
		}
		data.CurrentHeight = uint(result.Result)
	case "Near":
		lastPayload := strings.NewReader("{\"jsonrpc\": \"2.0\",\"id\": 1,\"method\":\"block\",\"params\":{\"finality\": \"final\"}}")
		req, _ := http.NewRequest("POST", data.HttpEndpoint, lastPayload)
		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		result := struct {
			Id      int64           `json:"id"`
			JsonRpc string          `json:"jsonrpc"`
			Result  NearBlockResult `json:"result"`
		}{}
		body, _ := io.ReadAll(res.Body)
		err := json.Unmarshal(body, &result)
		if err != nil {
			logger.Errorf("near json unmarshal failed: %s", err)
			return data
		}
		data.CurrentHeight = uint(result.Result.Header.Height)
		prePayload := strings.NewReader(fmt.Sprintf("{\"jsonrpc\": \"2.0\",\"id\": 1,\"method\":\"block\",\"params\":{\"block_id\": %d}}", result.Result.Header.PrevHeight))
		req, _ = http.NewRequest("POST", data.HttpEndpoint, prePayload)
		res, _ = http.DefaultClient.Do(req)
		defer res.Body.Close()
		resultData := struct {
			Id      int64           `json:"id"`
			JsonRpc string          `json:"jsonrpc"`
			Result  NearBlockResult `json:"result"`
		}{}
		body, _ = io.ReadAll(res.Body)
		err = json.Unmarshal(body, &resultData)
		if err != nil {
			logger.Errorf("near get pre block failed: %s", err)
			return data
		}
		preSeconds := resultData.Result.Header.Timestamp / 1e9     // 获取秒级时间戳
		preNanoseconds := resultData.Result.Header.Timestamp % 1e9 // 获取纳秒部分
		lastSeconds := result.Result.Header.Timestamp / 1e9
		lastNanoseconds := result.Result.Header.Timestamp % 1e9
		blockTime := time.Unix(lastSeconds, lastNanoseconds).Sub(time.Unix(preSeconds, preNanoseconds))
		data.BlockTime = blockTime.String()
		return data
	}
	return data
}

type NearBlockResult struct {
	Header NearBlockHead `json:"header"`
}
type NearBlockHead struct {
	Height     int64 `json:"height"`
	Timestamp  int64 `json:"timestamp"`
	PrevHeight int64 `json:"prev_height"`
}

func QueryBlockTime(data node.NodeDetail) string {
	switch string(data.ChainProtocol) {
	case "Ethereum", "Optimism", "Avalanche":
		client, err := rpc.Dial(data.HttpEndpoint)
		if err != nil {
			logger.Errorf("get client failed: %s", err)
			return ""
		}
		var lastNumber string
		err = client.Call(&lastNumber, "eth_blockNumber")
		if err != nil {
			logger.Errorf("get last block number failed: %s", err)
			return ""
		}
		latestBlockNumber := new(big.Int)
		latestBlockNumber.SetString(lastNumber[2:], 16)
		var latestBlock map[string]interface{}
		err = client.Call(&latestBlock, "eth_getBlockByNumber", lastNumber, true)
		if err != nil {
			logger.Errorf("get block by number failed: %s", err)
			return ""
		}
		latestBlockTimestamp := latestBlock["timestamp"].(string)
		latestBlockTime, err := strconv.ParseInt(latestBlockTimestamp, 0, 64)
		if err != nil {
			logger.Errorf("get last block time failed: %s", err)
			return ""
		}
		previousBlockNumber := new(big.Int).Sub(latestBlockNumber, big.NewInt(1))

		var previousBlock map[string]interface{}
		err = client.Call(&previousBlock, "eth_getBlockByNumber", fmt.Sprintf("0x%x", previousBlockNumber), true)
		if err != nil {
			logger.Errorf("get pre block number failed: %s", err)
			return ""
		}
		previousBlockTimestamp := previousBlock["timestamp"].(string)
		previousBlockTime, err := strconv.ParseInt(previousBlockTimestamp, 0, 64)
		if err != nil {
			logger.Errorf("get pre block time failed: %s", err)
			return ""
		}
		blockTime := time.Unix(latestBlockTime, 0).Sub(time.Unix(previousBlockTime, 0))
		return blockTime.String()
	}
	return ""
}
