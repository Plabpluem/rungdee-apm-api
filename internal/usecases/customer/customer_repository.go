package usecases

import (
	"rungdee-apm-api/internal/adapters/customer/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/customer/dto"
)

type CustomerRepository interface {
	Create(dto *dto.CreateCustomerDto) (*entities.Customer, error)
	Findall(dto *dto.FilterCustomerDto) (*response.CustomerPaginatedResponse, error)
	Find(dto *dto.FindCustomerDto) (*entities.Customer, error)
	Update(dto *dto.UpdateCustomerDto) (*entities.Customer, error)
	FindallDropdown() ([]*entities.Customer, error)
	GeneratePrescreen(dto *dto.CreateCustomerPrescreenDto) (*entities.CustomerLinePrescreen, error)
	UpdateLinePrescreen(dto *dto.UpdateLinePrescreenDto) (*entities.Customer, error)
}
