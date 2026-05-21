package dto

import "github.com/google/uuid"

type CreateInvoiceDto struct {
	ContractUuid uuid.UUID `json:"contract_uuid" validate:"required"`

	CurWaterUnit *float64 `json:"cur_water_unit"`

	CurElecUnit    float64 `json:"cur_elec_unit" validate:"required"`
	IsUnlimitWater *bool   `json:"is_unlimit_water"`
}
