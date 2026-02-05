package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleUser string

const (
	ADMIN    RoleUser = "admin"
	EMPLOYEE RoleUser = "employee"
)

type User struct {
	gorm.Model
	Uuid     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Role     RoleUser  `json:"role"`
}
