package dto

import "github.com/google/uuid"

type UpdateInvoiceDto struct {
	Uuid       uuid.UUID `json:"uuid"`
	ContractId uint      `json:"contract_id"`
	// RentPrice     float64   `json:"rent_price"`
	PrevWaterUnit float64  `json:"prev_water_unit"`
	CurWaterUnit  *float64 `json:"cur_water_unit"`
	// WaterUnit     float64   `json:"water_unit"`
	// WaterPrice    float64   `json:"water_price"`
	PrevElecUnit   float64 `json:"prev_elec_unit"`
	CurElecUnit    float64 `json:"cur_elec_unit"`
	IsUnlimitWater *bool   `json:"is_unlimit_water"`

	// ElecUnit      float64   `json:"elec_unit"`
	// ElecPrice     float64   `json:"elec_price"`
	// TotalAmount   float64   `json:"total_amount"`
}
