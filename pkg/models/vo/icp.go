package vo

type UserIcpInfoVo struct {
	UserId     int    `json:"userId"`
	AccountId  string `json:"accountId"`
	IcpBalance string `json:"icpBalance"`
}

type IcpAccountVo struct {
	UserId        int  `json:"userId"`
	AccountIdFlag bool `json:"accountIdFlag"`
	WalletIdFlag  bool `json:"walletIdFlag"`
}

type IcpCanisterBalanceVo struct {
	UserId        int    `json:"userId"`
	CanisterId    string `json:"canisterId"`
	CyclesBalance string `json:"cyclesBalance"`
}

type IcpCanisterPage struct {
	Data     []IcpCanisterVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type CanisterStatusRes struct {
	Status  string `json:"status"`
	Balance string `json:"balance"`
}

// New
type AccountBrief struct {
	Canisters int `json:"canisters"`
	Running   int `json:"running"`
	Stopped   int `json:"stopped"`
}

type AccountOverview struct {
	Canisters int    `json:"canisters"`
	Projects  int    `json:"projects"`
	Cycles    string `json:"cycles"`
	Icps      string `json:"icps"`
}

type UserCanisterVo struct {
	CanisterId   string `json:"canisterId"`
	CanisterName string `json:"canisterName"`
	Cycles       string `json:"cycles"`
	Status       string `json:"status"`
	Project      string `json:"project"`
	UpdateAt     string `json:"updateAt"`
}

type UserCanisterPage struct {
	Data     []UserCanisterVo `json:"data"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
}

type CanisterOverview struct {
	CanisterId   string `json:"canisterId"`
	CanisterName string `json:"canisterName"`
	Project      string `json:"project"`
	Cycles       string `json:"cycles"`
	Status       string `json:"status"`
	Size         string `json:"size"`
	ModuleHash   string `json:"moduleHash"`
	UpdateAt     string `json:"updateAt"`
}

type ControllerVo struct {
	PrincipalId string `json:"principalId"`
	Scope       string `json:"scope"`
	Type        string `json:"type"`
}

type ControllerPage struct {
	Data     []ControllerVo `json:"data"`
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}
