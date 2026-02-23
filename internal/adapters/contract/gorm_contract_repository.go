package adapters

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/adapters/contract/response"
	"rungdee-apm-api/internal/entities"

	// usecases "rungdee-apm-api/internal/usecases/contract"
	"rungdee-apm-api/internal/usecases/contract/dto"
	"rungdee-apm-api/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormContract struct {
	db *gorm.DB
}

func NewGormContractRepository(db *gorm.DB) *GormContract {
	return &GormContract{db: db}
}

func (r *GormContract) Create(dto *dto.CreateContractDto) (*entities.Contract, error) {
	contract := entities.Contract{
		RoomId:     dto.RoomId,
		CustomerId: dto.CustomerId,
		StartDate:  dto.StartDate,
		EndDate:    dto.EndDate,
		Status:     dto.Status,
	}

	if err := r.db.Create(&contract).Error; err != nil {
		return nil, err
	}
	return &contract, nil

}

func (r *GormContract) Findall(dto *dto.FilterContractDto) (*response.ContractPaginatedResponse, error) {
	var contracts []*entities.Contract
	var total int64

	pagination := pkg.Pagination{
		Page:    dto.Page,
		PerPage: dto.PerPage,
	}

	db := r.db.Model(&entities.Contract{})
	db.Count(&total)

	err := db.Limit(pagination.GetPerPage()).Offset(pagination.GetOffSet()).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return &response.ContractPaginatedResponse{
		Data:       contracts,
		Total:      total,
		Page:       pagination.GetPage(),
		PerPage:    pagination.GetPerPage(),
		TotalPages: int((total + int64(pagination.GetPerPage()) - 1) / int64(pagination.GetPerPage())),
	}, nil

}

func (r *GormContract) Findone(dto *dto.FindContractDto) (*entities.Contract, error) {
	var contract entities.Contract

	err := r.db.Where("uuid = ?", dto.Uuid).Preload("Room").First(&contract).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found contract")
		}
		return nil, err
	}
	return &contract, err
}
func (r *GormContract) Update(dto *dto.UpdateContractDto) (*entities.Contract, error) {
	var contract entities.Contract

	err := r.db.Where("uuid = ?", dto.Uuid).First(&contract).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found contract")
		}
		return nil, err
	}

	err = r.db.Model(&contract).Updates(&entities.Contract{
		RoomId:     dto.RoomId,
		CustomerId: dto.CustomerId,
		StartDate:  dto.StartDate,
		EndDate:    dto.EndDate,
		Status:     dto.Status,
	}).Error

	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (r *GormContract) FindByUuid(uuid uuid.UUID) (*entities.Contract, error) {
	var contract entities.Contract

	err := r.db.Where("uuid = ?", uuid).Preload("Room").Preload("Invoice").First(&contract).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found contract")
		}
		return nil, err
	}
	return &contract, err
}

func (r *GormContract) FindById(id uint) (*entities.Contract, error) {
	var contract entities.Contract

	err := r.db.Where("id = ?", id).Preload("Room").Preload("Invoice").First(&contract).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found contract")
		}
		return nil, err
	}
	return &contract, err
}
