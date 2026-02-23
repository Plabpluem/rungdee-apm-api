package dto

import "github.com/google/uuid"

type FindInvoiceDto struct {
	UUid uuid.UUID `json:"uuid"`
}
