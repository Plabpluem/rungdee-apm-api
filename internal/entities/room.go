package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Uuid         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Number       string    `json:"number"`
	RentPrice    float64   `json:"rent_price"`
	WaterPerUnit float64   `json:"water_per_unit"`
	ElecPerUnit  float64   `json:"elec_per_unit"`

	Contract []Contract `gorm:"foreignKey:RoomId" json:"contract"`
}
