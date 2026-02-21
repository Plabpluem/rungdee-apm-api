package dto

import "github.com/google/uuid"

type UpdateCustomerDto struct {
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	LastName   string    `json:"last_name"`
	IdCard     string    `json:"id_card"`
	LineUserId string    `json:"line_user_id"`
	IsActive   bool      `json:"is_active"`
}
