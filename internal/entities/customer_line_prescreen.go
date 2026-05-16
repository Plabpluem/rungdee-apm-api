package entities

import (
	"time"

	"gorm.io/gorm"
)

type CustomerLinePrescreen struct {
	gorm.Model
	Ref        string     `json:"ref"`
	CustomerId uint       `json:"customer_id"`
	ExpireDate time.Time  `json:"expire_date"`
	DeleteAt   *time.Time `json:"delete_at"`
}
