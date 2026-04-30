package dto

type CreateCustomerPrescreenDto struct {
	CustomerId uint   `json:"customer_id" validate:"required"`
	Ref        string `json:"ref"`
}
