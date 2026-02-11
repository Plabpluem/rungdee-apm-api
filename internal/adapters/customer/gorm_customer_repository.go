package adapter

import (
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/customer"

	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) usecases.CustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) Create(dto *entities.Customer) (*entities.Customer, error) {
	err := r.db.Create(dto).Error

	if err != nil {
		return nil, err
	}
	return dto, err
}
