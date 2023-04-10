package models

import "time"

type OracleRequestEvent struct {
	TransactionHash    string    `json:"transaction_hash"`
	RequestingContract string    `json:"requesting_contract"`
	RequestInitiator   string    `json:"request_initiator"`
	SubscriptionId     uint64    `json:"subscription_id"`
	SubscriptionOwner  string    `json:"subscription_owner"`
	BlockNumber        uint64    `json:"block_number"`
	TxIndex            uint      `json:"tx_index"`
	BlockHash          string    `json:"block_hash"`
	Index              uint      `json:"index" gorm:"column:_index"`
	Removed            bool      `json:"removed"`
	Chain              string    `json:"chain"`
	Network            string    `json:"network"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type OracleRequestEventAndName struct {
	Name string `json:"name"`
	OracleRequestEvent
}
