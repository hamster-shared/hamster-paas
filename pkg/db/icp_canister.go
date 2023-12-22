package db

import "database/sql"

type CanisterStatus int

const (
	Processing CanisterStatus = iota + 1
	Running
	Stopped
)

func (s *CanisterStatus) String() string {
	switch *s {
	case Processing:
		return "Processing"
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	default:
		return "Unknown"
	}
}

func DBStatus(st string) CanisterStatus {
	switch st {
	case "Running":
		return Running
	case "Stopped":
		return Stopped
	default:
		return Processing
	}
}

type IcpCanister struct {
	Id           int            `json:"id"`
	FkUserId     uint           `json:"fkUserId"`
	ProjectId    string         `json:"projectId"`
	CanisterId   string         `json:"canisterId"`
	CanisterName string         `json:"canisterName"`
	Cycles       sql.NullString `gorm:"type:decimal(10,2)" json:"cycles"`
	Status       CanisterStatus `json:"status"`
	CreateTime   sql.NullTime   `json:"createTime"`
	UpdateTime   sql.NullTime   `json:"updateTime"`
	Contract     string         `json:"contract"`
}
