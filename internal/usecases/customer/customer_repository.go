package usecases

import "rungdee-apm-api/internal/entities"

type CustomerRepository interface {
	Create(dto *entities.Customer) (*entities.Customer, error)
}
