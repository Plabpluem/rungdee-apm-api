package adapter

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/adapters/customer/response"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/customer"
	"rungdee-apm-api/internal/usecases/customer/dto"
	"rungdee-apm-api/pkg"

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

func (r *GormCustomerRepository) Findall(dto *dto.FilterCustomerDto) (*response.CustomerPaginatedResponse, error) {

	var customer []*entities.Customer
	var total int64

	pagination := pkg.Pagination{
		Page:    dto.Page,
		PerPage: dto.PerPage,
	}

	db := r.db.Model(&entities.Customer{})
	db.Count(&total)

	err := db.Limit(pagination.GetPerPage()).Offset(pagination.GetOffSet()).Find(&customer).Error
	if err != nil {
		return nil, err
	}
	return &response.CustomerPaginatedResponse{
		Data:       customer,
		Total:      total,
		Page:       pagination.GetPage(),
		PerPage:    pagination.GetPerPage(),
		TotalPages: int((total + int64(pagination.GetPerPage()) - 1) / int64(pagination.GetPerPage())),
	}, nil
}

func (r *GormCustomerRepository) Find(dto *dto.FindCustomerDto) (*entities.Customer, error) {
	var customer entities.Customer

	err := r.db.Where("uuid = ?", dto.UUid).First(&customer).Error

	if err != nil {
		return nil, err
	}

	return &customer, nil
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

	return &customer, nil
}
