package db

import "database/sql"

type IcpConsumption struct {
	Id         int            `json:"id"`
	CanisterId string         `json:"canisterId"`
	Cycles     sql.NullString `gorm:"type:decimal(10,2)" json:"cycles"`
	UpdateTime sql.NullTime   `json:"updateTime"`
}
