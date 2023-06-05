package models

import "time"

type FunctionConsumerEvent struct {
	Id        int64     `json:"id"`
	RequestId string    `json:"requestId"`
	Result    string    `json:"result"`
	ErrorInfo string    `json:"errorInfo"`
	Created   time.Time `json:"created"`
}
