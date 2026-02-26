package dto

import "github.com/google/uuid"

type CreateInvoiceDto struct {
	ContractUuid uuid.UUID `json:"contract_uuid" validate:"required"`

	CurWaterUnit float64 `json:"cur_water_unit" validate:"required"`

	CurElecUnit float64 `json:"cur_elec_unit" validate:"required"`
}
