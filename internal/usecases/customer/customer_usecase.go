package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/customer/dto"
)

// ตั้งชื่อ infaceface ของ usecase
type CustomerUseCase interface {
	Create(dto *dto.CreateCustomerDto) (*entities.Customer, error)
	Findall() ([]*entities.Customer, error)
	Find(dto *dto.FindCustomerDto) (*entities.Customer, error)
	Update(dto *dto.UpdateCustomerDto) (*entities.Customer, error)
}

func NewCustomerService(repo CustomerRepository) CustomerUseCase {
	return &CustomerService{repo: repo}
}

type CustomerService struct {
	repo CustomerRepository
}

func (s *CustomerService) Create(dto *dto.CreateCustomerDto) (*entities.Customer, error) {
	return s.repo.Create(dto)
}

func (s *CustomerService) Findall() ([]*entities.Customer, error) {
	return s.repo.Findall()
}
func (s *CustomerService) Find(dto *dto.FindCustomerDto) (*entities.Customer, error) {
	return s.repo.Find(dto)
}
func (s *CustomerService) Update(dto *dto.UpdateCustomerDto) (*entities.Customer, error) {
	return s.repo.Update(dto)
}
