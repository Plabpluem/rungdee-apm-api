package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/customer/dto"
)

type CustomerRepository interface {
	Create(dto *dto.CreateCustomerDto) (*entities.Customer, error)
	Findall() ([]*entities.Customer, error)
	Find(dto *dto.FindCustomerDto) (*entities.Customer, error)
	Update(dto *dto.UpdateCustomerDto) (*entities.Customer, error)
}
