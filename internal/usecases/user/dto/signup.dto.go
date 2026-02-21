package dto

import "rungdee-apm-api/internal/entities"

type SignUpDto struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Name     string            `json:"name"`
	Role     entities.RoleUser `json:"role"`
}
