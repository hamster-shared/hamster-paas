package vo

import (
	uuid "github.com/iris-contrib/go.uuid"
)

type AlineProjectIDAndName struct {
	Id     uuid.UUID `gorm:"primaryKey" json:"id"`
	Name   string    `json:"name"`
	UserId int64     `json:"UserId"`
}
