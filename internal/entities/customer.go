package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Uuid       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	IdCard     string    `json:"id_card"`
	LineUserId string    `json:"line_user_id"`
	IsActive   bool      `json:"is_active"`

	Contract *Contract `gorm:"foreignKey:CustomerId" json:"contract"`
}
