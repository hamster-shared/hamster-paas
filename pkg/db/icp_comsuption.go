package db

import "database/sql"

type IcpComsuption struct {
	CanisterId string         `json:"canisterId"`
	Cycles     sql.NullString `gorm:"type:decimal(10,2)" json:"cycles"`
	UpdateTime sql.NullTime   `json:"updateTime"`
}
