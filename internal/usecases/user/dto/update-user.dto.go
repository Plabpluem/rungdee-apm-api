package dto

import (
	"rungdee-apm-api/internal/entities"

	"github.com/google/uuid"
)

type UpdateUserDto struct {
	Uuid     uuid.UUID         `json:"uuid"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Name     string            `json:"name"`
	Role     entities.RoleUser `json:"role"`
}
