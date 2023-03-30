package vo

type ApiRequestRpcCreateApp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Chain       string `json:"chain"`
	Network     string `json:"network"`
	Account     string `json:"account"`
}
