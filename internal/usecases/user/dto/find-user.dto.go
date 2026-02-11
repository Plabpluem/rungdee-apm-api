package dto

import (
	"github.com/google/uuid"
)

type FindUserDto struct {
	Uuid uuid.UUID `json:"uuid"`
}
