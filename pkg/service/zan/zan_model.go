package zan

type BaseResponse[T interface{}] struct {
	Code    *string `json:"code"`
	Data    T       `json:"data"`
	Message *string `json:"message"`
	Success bool    `json:"success"`
}

type PageResponse[T interface{}] struct {
	Data      []T   `json:"data"`
	PageCount int   `json:"pageCount"`
	PageNum   int   `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	Total     int64 `json:"total"`
}

type AuthUrl struct {
	AuthUrl string `json:"authUrl"`
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  string `json:"expireTime"`
}

type ApiKeyCreateReq struct {
	// API Key 名字
	Name string `json:"name"`
}

type ApiKeyBase struct {
	// API KEY ID
	ApiKeyId string `json:"apiKeyId"`
}

type ApiKeyDigestInfo struct {
	ApiKeyBase
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// API Key 名字
	Name string `json:"name"`
}

type ApiKeyDetailInfo struct {
	ApiKeyBase
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// API KEY的⽣态访问信息
	EcosystemDetailInfos []ApiKeyEcosystemDetail `json:"ecosystemDetailInfos"`
}

type ApiKeyEcosystemDetail struct {
	// ⽣态编码
	EcosystemCode string `json:"ecosystemCode"`
	// ⽣态icon地址
	EcosystemIcon string `json:"ecosystemIcon"`
	// ⽣态名
	EcosystemName string `json:"ecosystemName"`
	// 链⽹络信息
	NetworkDetailInfoList []NetworkDetailInfo `json:"networkDetailInfoList"`
}

type NetworkDetailInfo struct {
	// ⽹络编码
	Code string `json:"code"`
	// ⽹络https请求地址
	HttpsUrl string `json:"httpsUrl"`
	// ⽹络名称
	Name string `json:"name"`
	// ⽹络wss请求地址
	WssUrl string `json:"wssUrl"`
}

type StatCreditCostItem struct {
	// 数据时间时间戳
	DataTime int64 `json:"dataTime"`
	// credit消耗量
	TotalCredit int64 `json:"totalCredit"`
}

type StatMethodCountItem struct {
	// 数据时间时间戳
	DataTime int64 `json:"dataTime"`
	// rpc 请求的⽅法
	Method string `json:"method"`
	// 调用次数
	Num int32 `json:"num"`
}

type StatMethodRequestActivityDetail struct {
	// rpc 请求的⽅法
	Method string `json:"method"`
	// 成功次数
	TotalNum int32 `json:"totalNum"`
	// 失败次数
	FailedNum int32 `json:"failedNum"`
}

type StatCreditCostOrigin struct {
	// 数据时间时间戳
	DataTime int64 `json:"dataTime"`
	// 请求来源Ip
	OriginIp string `json:"originIp"`
	// 调⽤总数
	TotalNum int64 `json:"totalNum"`
	// http请求数量
	HttpsNum int64 `json:"httpsNum"`
	// wss请求数量
	WssNum int64 `json:"wssNum"`
}

type EcosystemDigestInfo struct {
	// ⽣态编码
	EcosystemCode string `json:"ecosystemCode"`
	// ⽣态名
	EcosystemName string `json:"ecosystemName"`
	// ⽣态icon地址
	EcosystemIcon string `json:"ecosystemIcon"`
	// 链⽹络摘要信息
	Networks []string `json:"networks"`
}

type PlanDetailInfo struct {
	// 套餐名
	PlanName string `json:"planName"`
	// 已使⽤credit数量
	UsedCredit int64 `json:"usedCredit"`
	// credit套餐总量
	TotalCredit int64 `json:"totalCredit"`
	// 加油包数量
	AdditionalCredit int64 `json:"additionalCredit"`
	// 过去24⼩时credit⽤量
	UsedCreditLast24Hours int64 `json:"usedCreditLast24Hours"`
	// credits限速（每秒）
	CreditLimit int64 `json:"creditLimit"`
	// 可开通API Key数量
	ApiKeyAmount int32 `json:"apiKeyAmount"`
	//  是 已开通API Key数量
	CreatedApiKeyAmount int32 `json:"createdApiKeyAmount"`
	// ⽣效时间
	ActiveTime string `json:"activeTime"`
	// 过期时间
	ExpireTime string `json:"expireTime"`
	// 刷新时间
	RefreshTime string `json:"refreshTime"`
	// 套餐状态, 套餐状态包含如下枚举：VALID、INVALID。
	Status string `json:"status"`
}
