package adapter

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/adapters/invoice/response"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/invoice"
	"rungdee-apm-api/internal/usecases/invoice/dto"
	"rungdee-apm-api/pkg"

	"gorm.io/gorm"
)

type GormInvoiceRepository struct {
	db *gorm.DB
}

func NewGormInvoiceRepository(db *gorm.DB) usecases.InvoiceRepository {
	return &GormInvoiceRepository{db: db}
}

func (r *GormInvoiceRepository) Create(dto *entities.Invoice) (*entities.Invoice, error) {
	err := r.db.Create(dto).Error

	if err != nil {
		return nil, err
	}
	return dto, nil
}

func (r *GormInvoiceRepository) Findall(dto *dto.FilterInvoiceDto) (*response.InvoicePaginatedResponse, error) {
	var invoice []*entities.Invoice
	var total int64

	pagination := pkg.Pagination{
		Page:    dto.Page,
		PerPage: dto.PerPage,
	}

	db := r.db.Model(&entities.Invoice{})
	db.Count(&total)

	err := db.Limit(pagination.GetPerPage()).Offset(pagination.GetOffSet()).Preload("Contract.Customer").Find(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &response.InvoicePaginatedResponse{
		Data:       invoice,
		Total:      total,
		Page:       pagination.GetPage(),
		PerPage:    pagination.GetPerPage(),
		TotalPages: int((total + int64(pagination.GetPerPage()) - 1) / int64(pagination.GetPerPage())),
	}, nil
}

func (r *GormInvoiceRepository) Find(dto *dto.FindInvoiceDto) (*entities.Invoice, error) {
	var invoice entities.Invoice

	err := r.db.Where("uuid = ?", dto.UUid).Preload("Contract.Room").Preload("Contract.Customer").First(&invoice).Error

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (r *GormInvoiceRepository) Update(dto *entities.Invoice) (*entities.Invoice, error) {
	var invoice entities.Invoice

	err := r.db.Where("uuid = ?", dto.Uuid).First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Not found invoice")
		}
		return nil, err
	}

	err = r.db.Model(&invoice).Updates(&entities.Invoice{
		ContractId:    dto.ContractId,
		RentPrice:     dto.RentPrice,
		WaterPrice:    dto.WaterPrice,
		ElecPrice:     dto.ElecPrice,
		WaterUnit:     dto.WaterUnit,
		PrevWaterUnit: dto.PrevWaterUnit,
		CurWaterUnit:  dto.CurWaterUnit,

		ElecUnit:       dto.ElecUnit,
		PrevElecUnit:   dto.PrevElecUnit,
		CurElecUnit:    dto.CurElecUnit,
		TotalAmount:    dto.TotalAmount,
		LinkPdf:        dto.LinkPdf,
		IsUnlimitWater: dto.IsUnlimitWater,
	}).Error

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
