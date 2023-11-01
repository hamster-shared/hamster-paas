package zan

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/samber/lo"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"
)

type ZanClient struct {
	baseUrl         string
	baseAccessToken string
	clientId        string
	privateKey      *rsa.PrivateKey
	httpclient      *http.Client
}

func NewZanClient(baseUrl string, accessToken string, clientId string, privateKeyPath string) *ZanClient {
	httpclient := http.DefaultClient
	httpclient.Timeout = time.Minute * 5

	// 读取私钥文件内容
	privateKeyPEM, err := os.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("无法读取私钥文件:", err)
		os.Exit(1)
	}

	// 解析 PEM 编码的私钥数据
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		fmt.Println("无效的 PEM 数据")
		os.Exit(1)
	}

	// 解析私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("私钥解析失败:", err)
		os.Exit(1)
	}

	return &ZanClient{
		baseUrl:         baseUrl,
		baseAccessToken: accessToken,
		clientId:        clientId,
		privateKey:      privateKey.(*rsa.PrivateKey),
		httpclient:      httpclient,
	}
}

// AuthUrl 1.1 获取授权链接接⼝
func (c *ZanClient) AuthUrl() (BaseResponse[AuthUrl], error) {
	params := map[string]string{}
	response, err := DoGet[AuthUrl](c, c.baseAccessToken, "/openapi/v1/account/auth-url", params)
	return response, err
}

// AccessToken 1.2 Oauth获取accessToken接⼝
func (c *ZanClient) AccessToken(authCode string) (BaseResponse[AccessToken], error) {
	params := map[string]string{
		"authCode": authCode,
	}

	response, err := DoGet[AccessToken](c, c.baseAccessToken, "/openapi/v1/account/oauth/access-token", params)
	return response, err
}

// ApiKeyCreate 2.1 创建API KEY接⼝
func (c *ZanClient) ApiKeyCreate(name string, accessToken string) (BaseResponse[ApiKeyBase], error) {

	req := ApiKeyCreateReq{
		Name: name,
	}

	response, err := DoPost[ApiKeyBase](c, accessToken, "/openapi/v1/node-service/api-keys", req)
	return response, err
}

// ApiKeyList 2.2 API KEY分⻚查询接⼝
func (c *ZanClient) ApiKeyList(page int, size int, accessToken string) (BaseResponse[PageResponse[ApiKeyDigestInfo]], error) {
	params := map[string]string{
		"page": strconv.Itoa(page),
		"size": strconv.Itoa(size),
	}
	response, err := DoGet[PageResponse[ApiKeyDigestInfo]](c, accessToken, "/openapi/v1/node-service/api-keys/list", params)

	return response, err
}

// ApiKeyDetail 2.3 API KEY详情接⼝
func (c *ZanClient) ApiKeyDetail(apiKeyId string, accessToken string) (BaseResponse[ApiKeyDetailInfo], error) {
	params := map[string]string{
		"apiKeyId": apiKeyId,
	}

	response, err := DoGet[ApiKeyDetailInfo](c, accessToken, "/openapi/v1/node-service/api-keys/detail", params)

	return response, err
}

// ApiKeyCreditCost 2.4 API KEY credit cost统计查询接⼝
func (c *ZanClient) ApiKeyCreditCost(apiKeyId string, accessToken string) (BaseResponse[[]StatCreditCostItem], error) {
	params := map[string]string{
		"apiKeyId": apiKeyId,
	}

	response, err := DoGet[[]StatCreditCostItem](c, accessToken, "/openapi/v1/node-service/api-keys/stats/credit-cost", params)

	return response, err
}

// ApiKeyRequestStats 2.5 API Key request统计查询接⼝
func (c *ZanClient) ApiKeyRequestStats(apiKeyId string, timeInterval string, ecosystem string, accessToken string) (BaseResponse[[]StatMethodCountItem], error) {
	params := map[string]string{
		"apiKeyId":     apiKeyId,
		"timeInterval": timeInterval,
		"ecosystem":    ecosystem,
	}

	response, err := DoGet[[]StatMethodCountItem](c, accessToken, "/openapi/v1/node-service/api-keys/stats/requests", params)
	return response, err
}

// ApiKeyRequestActivity 2.6 API KEY requests activity统计查询接⼝
func (c *ZanClient) ApiKeyRequestActivityStats(apiKeyId string, timeInterval string, ecosystem string, accessToken string) (BaseResponse[[]StatMethodRequestActivityDetail], error) {
	params := map[string]string{
		"apiKeyId":     apiKeyId,
		"timeInterval": timeInterval,
		"ecosystem":    ecosystem,
	}

	response, err := DoGet[[]StatMethodRequestActivityDetail](c, accessToken, "/openapi/v1/node-service/api-keys/stats/requests-activity", params)
	return response, err
}

// ApiKeyRequestOrigin 2.7 API KEY request Origin统计查询接⼝
func (c *ZanClient) ApiKeyRequestOriginStats(apiKeyId string, timeInterval string, accessToken string) (BaseResponse[[]StatCreditCostOrigin], error) {
	params := map[string]string{
		"apiKeyId":     apiKeyId,
		"timeInterval": timeInterval,
	}

	response, err := DoGet[[]StatCreditCostOrigin](c, accessToken, "/openapi/v1/node-service/api-keys/stats/requests-origin", params)
	return response, err
}

// EcosystemsDigest 2.8 链⽣态摘要信息查询接
func (c *ZanClient) EcosystemsDigest() (BaseResponse[[]EcosystemDigestInfo], error) {
	params := map[string]string{}

	response, err := DoGet[[]EcosystemDigestInfo](c, c.baseAccessToken, "/openapi/v1/node-service/ecosystems/digest", params)
	return response, err
}

// 2.9 套餐信息查询接⼝
func (c *ZanClient) Plan(accessToken string) (BaseResponse[PlanDetailInfo], error) {
	params := map[string]string{}

	response, err := DoGet[PlanDetailInfo](c, accessToken, "/openapi/v1/node-service/plan", params)
	return response, err
}

// 2.10 API KEY requests activity failed 统计查询接口
func (c *ZanClient) ApiKeyRequestActivityStatsFail(accessToken string, apiKeyId string, timeInterval string, ecosystem string, method string) (BaseResponse[[]StatMethodRequestActivityFailedDetailGwInfo], error) {
	params := map[string]string{
		"apiKeyId":     apiKeyId,
		"timeInterval": timeInterval,
		"ecosystem":    ecosystem,
		"method":       method,
	}

	response, err := DoGet[[]StatMethodRequestActivityFailedDetailGwInfo](c, accessToken, "/openapi/v1/node-service/api-keys/stats/requests-activity/failed", params)
	return response, err
}

func (c *ZanClient) getParamsSignature(timestamp int64, params map[string][]string) string {
	unsignedPayload := fmt.Sprintf("%d|", timestamp)
	if len(params) > 0 {
		sortedJson := sortMapAndJson(params)
		unsignedPayload += sortedJson
	}

	fmt.Println("1.待加密payload：" + unsignedPayload)
	hash := sha256.Sum256([]byte(unsignedPayload))
	hash2 := sha256.Sum256(hash[:])
	signed, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, hash2[:])
	if err != nil {
		return ""
	}
	signatureStr := base64.StdEncoding.EncodeToString(signed)

	fmt.Println("3.签名：" + signatureStr)
	return signatureStr
}

func (c *ZanClient) getBodySignature(timestamp int64, params any) string {
	unsignedPayload := fmt.Sprintf("%d|%s", timestamp, objectToJson(params))

	fmt.Println("1.待加密payload：" + unsignedPayload)
	hash := sha256.Sum256([]byte(unsignedPayload))
	hash2 := sha256.Sum256(hash[:])
	signed, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, hash2[:])
	if err != nil {
		return ""
	}
	signatureStr := base64.StdEncoding.EncodeToString(signed)

	fmt.Println("3.签名：" + signatureStr)
	return signatureStr
}

func sortMapAndJson(params map[string][]string) string {
	// 提取 map 的键到切片
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 创建一个新的有序 map
	sortedMap := make(map[string][]string)

	// 将排序后的键值对填充到新 map 中
	for _, k := range keys {
		v := params[k]
		sortedMap[k] = v
	}

	return objectToJson(sortedMap)
}

func objectToJson(obj any) string {
	jsonData, _ := json.Marshal(obj)
	return string(jsonData)
}

func DoGet[T any](c *ZanClient, accessToken string, path string, params map[string]string) (response BaseResponse[T], err error) {
	timestamp := time.Now().UnixMilli()
	fullPath, err := url.JoinPath(c.baseUrl, path)
	reqUrl, err := url.ParseRequestURI(fullPath)
	signParams := lo.MapEntries(params, func(key string, value string) (string, []string) {
		return key, []string{value}
	})

	fullUrl := reqUrl.String()
	if len(params) > 0 {
		fullUrl += "?" + url.Values(signParams).Encode()
	}

	req, _ := http.NewRequest("GET", fullUrl, nil)
	req.Header.Set("Request-Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Set("X-Access-Token", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Authorization", fmt.Sprintf("%s_%s", c.clientId, c.getParamsSignature(timestamp, signParams)))

	fmt.Println("5.请求URL和Header: ")
	fmt.Println(req.Method, req.URL.String())
	for key, value := range req.Header {
		fmt.Println(key, ":", value[0])
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)

	fmt.Println("----响应----")
	fmt.Println(string(respData))
	err = json.Unmarshal(respData, &response)

	return
}

func DoPost[T any](c *ZanClient, accessToken string, path string, params any) (response BaseResponse[T], err error) {
	timestamp := time.Now().UnixMilli()
	fullPath, err := url.JoinPath(c.baseUrl, path)
	reqUrl, err := url.ParseRequestURI(fullPath)
	requestBody, err := json.Marshal(params)
	req, _ := http.NewRequest("POST", reqUrl.String(), bytes.NewBuffer(requestBody))
	req.Header.Set("Request-Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Set("X-Access-Token", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Authorization", fmt.Sprintf("%s_%s", c.clientId, c.getBodySignature(timestamp, params)))

	fmt.Println("5.请求URL和Header: ")
	fmt.Println(req.Method, req.URL.String())
	for key, value := range req.Header {
		fmt.Println(key, ":", value[0])
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)

	fmt.Println("----响应----")
	fmt.Println(string(respData))
	err = json.Unmarshal(respData, &response)

	return
}
