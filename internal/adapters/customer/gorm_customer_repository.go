package adapter

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/customer"
	"rungdee-apm-api/internal/usecases/customer/dto"

	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) usecases.CustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) Create(dto *dto.CreateCustomerDto) (*entities.Customer, error) {
	customer := entities.Customer{
		Name:       dto.Name,
		LastName:   dto.LastName,
		IdCard:     dto.IdCard,
		LineUserId: dto.LineUserId,
		IsActive:   dto.IsActive,
	}
	err := r.db.Create(&customer).Error

	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *GormCustomerRepository) Findall() ([]*entities.Customer, error) {
	var customer []*entities.Customer
	var total int64

	db := r.db.Model(&entities.Customer{})
	db.Count(&total)

	err := db.Find(&customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *GormCustomerRepository) Find(dto *dto.FindCustomerDto) (*entities.Customer, error) {
	var customer entities.Customer

	err := r.db.Where("uuid = ?", dto.UUid).First(&customer).Error

	if err != nil {
		return nil, err
	}

	return &customer, err
}

func (r *GormCustomerRepository) Update(dto *dto.UpdateCustomerDto) (*entities.Customer, error) {
	var customer entities.Customer

	err := r.db.Where("uuid = ?", dto.Uuid).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found customer")
		}
		return nil, err
	}

	err = r.db.Model(&customer).Updates(&entities.Customer{
		Name:       dto.Name,
		LastName:   dto.LastName,
		IdCard:     dto.IdCard,
		LineUserId: dto.LineUserId,
		IsActive:   dto.IsActive,
	}).Error

	if err != nil {
		return nil, err
	}

	return &customer, err
}
