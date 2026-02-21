package dto

import "github.com/google/uuid"

type FindCustomerDto struct {
	UUid uuid.UUID `json:"uuid"`
}
