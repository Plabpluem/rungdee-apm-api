package dto

type CreateCustomerDto struct {
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	IdCard     string `json:"id_card"`
	LineUserId string `json:"line_user_id"`
	IsActive   bool   `json:"is_active"`
}
