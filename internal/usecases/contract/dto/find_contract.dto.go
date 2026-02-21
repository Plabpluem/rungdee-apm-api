package dto

import "github.com/google/uuid"

type FindContractDto struct {
	Uuid uuid.UUID `json:"uuid"`
}
