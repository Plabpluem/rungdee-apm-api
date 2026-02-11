package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Uuid          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	ContractId    uint      `json:"contract_id"`
	RentPrice     float64   `json:"rent_price"`
	PrevWaterUnit float64   `json:"prev_water_unit"`
	CurWaterUnit  float64   `json:"cur_water_unit"`
	WaterUnit     float64   `json:"water_unit"`
	WaterPrice    float64   `json:"water_price"`
	PrevElecUnit  float64   `json:"prev_elec_unit"`
	CurElecUnit   float64   `json:"cur_elec_unit"`
	ElecUnit      float64   `json:"elec_unit"`
	ElecPrice     float64   `json:"elec_price"`
	TotalAmount   float64   `json:"total_amount"`

	Contract *Contract `gorm:"foreignKey:ContractId;references:ID" json:"contract"`
}
