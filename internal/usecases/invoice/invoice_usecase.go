package usecases

import (
	"bytes"
	"fmt"
	"rungdee-apm-api/internal/adapters/invoice/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/invoice/dto"
	usecases "rungdee-apm-api/internal/usecases/line"
	storage_usecases "rungdee-apm-api/internal/usecases/storage"
	"time"
)

type InvoiceUseCase interface {
	Create(req *dto.CreateInvoiceDto) (*entities.Invoice, error)
	Findall(dto *dto.FilterInvoiceDto) (*response.InvoicePaginatedResponse, error)
	Find(dto *dto.FindInvoiceDto) (*entities.Invoice, error)
	Update(req *dto.UpdateInvoiceDto) (*entities.Invoice, error)
	Generate(dto *dto.FindInvoiceDto) ([]byte, error)
	CreatePdf(req *dto.FindInvoiceDto) (*entities.StorageResponse, error)
}

func NewInvoiceService(repo InvoiceRepository, contractRepo ContractReader, pdfRepo PdfGenerate, lineRepo usecases.LineRepository, storageRepo storage_usecases.StorageRepository) InvoiceUseCase {
	return &InvoiceService{repo: repo, contractRepo: contractRepo, pdfRepo: pdfRepo, lineRepo: lineRepo, storageRepo: storageRepo}
}

type InvoiceService struct {
	repo         InvoiceRepository
	contractRepo ContractReader
	pdfRepo      PdfGenerate
	lineRepo     usecases.LineRepository
	storageRepo  storage_usecases.StorageRepository
}

func (s *InvoiceService) Create(req *dto.CreateInvoiceDto) (*entities.Invoice, error) {
	contract, err := s.contractRepo.FindByUuid(req.ContractUuid)
	if err != nil {
		return nil, err
	}

	var prev_elec_unit float64
	var prev_water_unit float64
	if len(*contract.Invoice) > 0 {
		lastItem := (*contract.Invoice)[len(*contract.Invoice)-1]
		prev_elec_unit = lastItem.CurElecUnit
		prev_water_unit = lastItem.CurWaterUnit
	} else {
		prev_elec_unit = contract.StartElecUnit
		prev_water_unit = contract.StartWaterUnit
	}

	elec_unit := req.CurElecUnit - prev_elec_unit
	water_unit := req.CurWaterUnit - prev_water_unit

	data_invoice := &entities.Invoice{
		ContractId:    contract.ID,
		RentPrice:     contract.Room.RentPrice,
		WaterPrice:    contract.Room.WaterPerUnit * water_unit,
		ElecPrice:     contract.Room.ElecPerUnit * elec_unit,
		ElecPerUnit:   contract.Room.ElecPerUnit,
		WaterPerUnit:  contract.Room.WaterPerUnit,
		WaterUnit:     water_unit,
		PrevWaterUnit: prev_water_unit,
		CurWaterUnit:  req.CurWaterUnit,

		ElecUnit:     elec_unit,
		PrevElecUnit: prev_elec_unit,
		CurElecUnit:  req.CurElecUnit,
		TotalAmount:  contract.Room.RentPrice + (contract.Room.WaterPerUnit * water_unit) + (contract.Room.ElecPerUnit * elec_unit),
	}
	invoice, err := s.repo.Create(data_invoice)

	invoice_dto := &dto.FindInvoiceDto{
		UUid: invoice.Uuid,
	}

	pdf_link, err := s.CreatePdf(invoice_dto)
	if err != nil {
		return nil, err
	}

	dto := &entities.Invoice{
		Uuid:    invoice.Uuid,
		LinkPdf: pdf_link.BaseUrl,
	}
	return s.repo.Update(dto)
}

func (s *InvoiceService) Findall(dto *dto.FilterInvoiceDto) (*response.InvoicePaginatedResponse, error) {
	return s.repo.Findall(dto)
}

func (s *InvoiceService) Find(dto *dto.FindInvoiceDto) (*entities.Invoice, error) {
	return s.repo.Find(dto)
}

func (s *InvoiceService) Update(req *dto.UpdateInvoiceDto) (*entities.Invoice, error) {
	contract, err := s.contractRepo.FindById(req.ContractId)
	if err != nil {
		return nil, err
	}

	var prev_elec_unit float64
	var prev_water_unit float64
	if len(*contract.Invoice)-1 > 0 {
		lastItem := (*contract.Invoice)[len(*contract.Invoice)-2]
		prev_elec_unit = lastItem.CurElecUnit
		prev_water_unit = lastItem.CurWaterUnit
	} else {
		prev_elec_unit = req.PrevElecUnit
		prev_water_unit = req.PrevWaterUnit
	}

	elec_unit := req.CurElecUnit - prev_elec_unit
	water_unit := req.CurWaterUnit - prev_water_unit

	dto := &entities.Invoice{
		Uuid:          req.Uuid,
		ContractId:    contract.ID,
		RentPrice:     contract.Room.RentPrice,
		WaterPrice:    contract.Room.WaterPerUnit * water_unit,
		ElecPrice:     contract.Room.ElecPerUnit * elec_unit,
		ElecPerUnit:   contract.Room.ElecPerUnit,
		WaterPerUnit:  contract.Room.WaterPerUnit,
		WaterUnit:     water_unit,
		PrevWaterUnit: prev_water_unit,
		CurWaterUnit:  req.CurWaterUnit,

		ElecUnit:     elec_unit,
		PrevElecUnit: prev_elec_unit,
		CurElecUnit:  req.CurElecUnit,
		TotalAmount:  contract.Room.RentPrice + (contract.Room.WaterPerUnit * water_unit) + (contract.Room.ElecPerUnit * elec_unit),
	}
	return s.repo.Update(dto)
}

func (s *InvoiceService) Generate(req *dto.FindInvoiceDto) ([]byte, error) {
	invoice, err := s.repo.Find(req)
	if err != nil {
		return nil, err
	}

	return s.pdfRepo.Generate(&PdfData{
		Invoice:  invoice,
		Room:     invoice.Contract.Room,
		Customer: invoice.Contract.Customer,
	})
}

func (s *InvoiceService) CreatePdf(req *dto.FindInvoiceDto) (*entities.StorageResponse, error) {
	invoice, err := s.repo.Find(req)
	if err != nil {
		return nil, err
	}

	generate_pdf, err := s.pdfRepo.Generate(&PdfData{
		Invoice:  invoice,
		Room:     invoice.Contract.Room,
		Customer: invoice.Contract.Customer,
	})
	if err != nil {
		return nil, err
	}

	uploadPdf, err := s.storageRepo.Save(bytes.NewReader(generate_pdf), fmt.Sprintf("invoice-%s-%s-%s.pdf", invoice.Contract.Room.Number, invoice.Contract.Customer.Name, time.Now().Format("dd-mm-YYYY")), "application/pdf")
	if err != nil {
		return nil, err
	}

	return uploadPdf, nil
}
