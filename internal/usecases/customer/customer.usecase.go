package usecases

import "rungdee-apm-api/internal/entities"

// ตั้งชื่อ infaceface ของ usecase
type CustomerUseCase interface {
	Create(dto *entities.Customer) (*entities.Customer, error)
}

func NewCustomerService(repo CustomerRepository) CustomerUseCase {
	return &CustomerService{repo: repo}
}

type CustomerService struct {
	repo CustomerRepository
}

func (s *CustomerService) Create(dto *entities.Customer) (*entities.Customer, error) {
	return s.repo.Create(dto)
}
