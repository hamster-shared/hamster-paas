package vo

type UserIcpInfoVo struct {
	UserId     int    `json:"userId"`
	AccountId  string `json:"accountId"`
	IcpBalance string `json:"icpBalance"`
}

type HasAccountVo struct {
	UserId       int  `json:"userId"`
	HasAccountId bool `json:"hasAccountId"`
	HasWalletId  bool `json:"hasWalletId"`
}

type IcpAccountVo struct {
	UserId      int    `json:"userId"`
	AccountId   string `json:"accountId"`
	PrincipalId string `json:"walletId"`
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

// New call Status
type CanisterStatus struct {
	Controllers []string `json:"controller"`
	Status      string   `json:"status"`
	Balance     string   `json:"balance"`
	MemorySize  string   `json:"memorySize"`
	ModuleHash  string   `json:"moduleHash"`
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
	MemorySize   string `json:"moduleSize"`
	ModuleHash   string `json:"moduleHash"`
	UpdateAt     string `json:"updateAt"`
}

// controllers
type ControllerVo struct {
	PrincipalId string `json:"principalId"`
	Type        string `json:"type"`
}

type ControllerPage struct {
	CanisterId string         `json:"canisterId" binding:"required"`
	Data       []ControllerVo `json:"data"`
	Total      int            `json:"total"`
	Page       int            `json:"page"`
	PageSize   int            `json:"pageSize"`
}

// Consumption
type ConsumptionVo struct {
	Cycles   string `json:"cycles"`
	UpdateAt string `json:"updateAt"`
}

type ConsumptionPage struct {
	Data     []ConsumptionVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type CreateCanisterParam struct {
	CanisterName string `json:"canisterName" binding:"required"`
}

type DeleteCanisterParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
}

type RechargeCanisterParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Amount     string `json:"amount" binding:"required"`
}

type AddControllerParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Controller string `json:"controller" binding:"required"`
}

type DelControllerParam struct {
	CanisterId string `json:"canisterId" binding:"required"`
	Controller string `json:"controller" binding:"required"`
}

type StatusType int

const (
	Running StatusType = iota + 1
	Stopped
)

type ChangeStatusParam struct {
	CanisterId string     `json:"canisterId" binding:"required"`
	Status     StatusType `json:"status" binding:"required"` // 1: running, 2: stopped
}

type InstallMode int

const (
	Install InstallMode = iota
	Upgrade
	Reinstall
)

func (i *InstallMode) String() string {
	return [...]string{"install", "upgrade", "reinstall"}[*i]
}

type InstallParam struct {
	CanisterId string      `json:"canisterId" binding:"required"`
	WasmType   string      `json:"wasmType" binding:"required"`
	Mode       InstallMode `json:"mode" binding:"required"` // 0: install 1: upgrade, 2: reinstall
}
