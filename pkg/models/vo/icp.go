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

type UserCycleInfoVo struct {
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

// user canisters
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

// controllers
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

// Consumption
type ConsumptionVo struct {
	Cycles     string `json:"cycles"`
	ModuleHash string `json:"moduleHash"`
	UpdateAt   string `json:"updateAt"`
}

type ConsumptionPage struct {
	Data     []ConsumptionVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type AddCanisterParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	ProjectId  string `json:"projectId" binding:"required"`
}

type RechargeCanisterParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
}

type AddControllerParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Controller string `json:"controller" binding:"required"`
	Role       int    `json:"role" binding:"required"` // 0: controller, 1: operator
}

type DelControllerParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Controller string `json:"controller" binding:"required"`
}

type ChangeStatusParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Status     int    `json:"status" binding:"required"` // 1: running, 2: stopped
}

type DeleteCanisterParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
}

type InstallDappParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	WasmFile   string `json:"wasmFile" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Mode       int    `json:"mode" binding:"required"` // 0: upgrade, 1: reinstall
}
