package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusContract string

const (
	ACTIVE   StatusContract = "active"
	INACTIVE StatusContract = "inactive"
)

type Contract struct {
	gorm.Model
	Uuid       uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	RoomId     uint           `json:"room_id"`
	CustomerId uint           `gorm:"uniqueIndex" json:"customer_id"`
	StartDate  time.Time      `gorm:"type:date" json:"start_date"`
	EndDate    time.Time      `gorm:"type:date" json:"end_date"`
	Status     StatusContract `gorm:"type:varchar(20);default:'active'" json:"status"`

	Room     Room       `gorm:"foreignKey:RoomId;references:ID" json:"room"`
	Customer *Customer  `gorm:"foreignKey:CustomerId;references:ID" json:"customer"`
	Invoice  *[]Invoice `gorm:"foreignKey:ContractId" json:"invoice"`
}
