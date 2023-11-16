package vo

type MiddleWareRpcZan struct {
	EcosystemCode string   `json:"ecosystemCode"`
	EcosystemName string   `json:"ecosystemName"`
	EcosystemIcon string   `json:"ecosystemIcon"`
	Networks      []string `json:"networks"`
	BuyFlag       bool     `json:"buyFlag"`
}
