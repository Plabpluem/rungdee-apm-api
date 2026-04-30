package adapter

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/adapters/customer/response"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/customer"
	"rungdee-apm-api/internal/usecases/customer/dto"
	"rungdee-apm-api/pkg"
	"time"

	"gorm.io/gorm"
)

type GormCustomerRepository struct {
	db *gorm.DB
}

func NewGormCustomerRepository(db *gorm.DB) usecases.CustomerRepository {
	return &GormCustomerRepository{db: db}
}

func (r *GormCustomerRepository) Create(dto *dto.CreateCustomerDto) (*entities.Customer, error) {
	isActive := true
	if dto.IsActive != nil {
		isActive = *dto.IsActive
	}
	customer := entities.Customer{
		Name:       dto.Name,
		LastName:   dto.LastName,
		IdCard:     dto.IdCard,
		LineUserId: dto.LineUserId,
		IsActive:   &isActive,
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

func (r *GormCustomerRepository) FindallDropdown() ([]*entities.Customer, error) {
	var customer []*entities.Customer
	db := r.db.Model(&entities.Customer{})

	err := db.Find(&customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *GormCustomerRepository) GeneratePrescreen(dto *dto.CreateCustomerPrescreenDto) (*entities.CustomerLinePrescreen, error) {
	customer_prescreen := entities.CustomerLinePrescreen{
		CustomerId: dto.CustomerId,
		ExpireDate: time.Now().Add(time.Hour * 1),
		Ref:        dto.Ref,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entities.CustomerLinePrescreen{}).
			Where("customer_id = ? AND delete_at IS NULL", dto.CustomerId).
			Update("delete_at", time.Now()).Error; err != nil {
			return err
		}

		if err := tx.Create(&customer_prescreen).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &customer_prescreen, nil
}

func (r *GormCustomerRepository) UpdateLinePrescreen(dto *dto.UpdateLinePrescreenDto) (*entities.Customer, error) {
	var customer entities.Customer
	var prescreen entities.CustomerLinePrescreen

	err := r.db.Where("ref = ? AND delete_at IS NULL", dto.Ref).First(&prescreen).Error
	if err != nil {
		return nil, fmt.Errorf("ไม่พบ ref นี้หรือ หมดอายุแล้ว")
	}

	if err = r.db.Model(&prescreen).Update("delete_at", time.Now()).Error; err != nil {
		return nil, err
	}

	err = r.db.Where("id = ?", prescreen.CustomerId).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found customer")
		}
		return nil, err
	}

	err = r.db.Model(&customer).Updates(&entities.Customer{
		LineUserId: dto.LineUserId,
	}).Error

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
