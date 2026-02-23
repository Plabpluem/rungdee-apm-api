package dto

import "github.com/google/uuid"

type CreateInvoiceDto struct {
	ContractUuid uuid.UUID `json:"contract_uuid"`

	PrevWaterUnit float64 `json:"prev_water_unit"`
	CurWaterUnit  float64 `json:"cur_water_unit"`

	PrevElecUnit float64 `json:"prev_elec_unit"`
	CurElecUnit  float64 `json:"cur_elec_unit"`
}
