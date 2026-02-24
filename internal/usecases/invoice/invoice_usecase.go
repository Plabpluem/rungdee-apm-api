package usecases

import (
	"rungdee-apm-api/internal/adapters/invoice/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/invoice/dto"
)

type InvoiceUseCase interface {
	Create(req *dto.CreateInvoiceDto) (*entities.Invoice, error)
	Findall(dto *dto.FilterInvoiceDto) (*response.InvoicePaginatedResponse, error)
	Find(dto *dto.FindInvoiceDto) (*entities.Invoice, error)
	Update(req *dto.UpdateInvoiceDto) (*entities.Invoice, error)
	Generate(dto *dto.FindInvoiceDto) ([]byte, error)
}

func NewInvoiceService(repo InvoiceRepository, contractRepo ContractReader, pdfRepo PdfGenerate) InvoiceUseCase {
	return &InvoiceService{repo: repo, contractRepo: contractRepo, pdfRepo: pdfRepo}
}

type InvoiceService struct {
	repo         InvoiceRepository
	contractRepo ContractReader
	pdfRepo      PdfGenerate
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
		prev_elec_unit = req.PrevElecUnit
		prev_water_unit = req.PrevWaterUnit
	}

	elec_unit := req.CurElecUnit - prev_elec_unit
	water_unit := req.CurWaterUnit - prev_water_unit

	dto := &entities.Invoice{
		ContractId:    contract.ID,
		RentPrice:     contract.Room.RentPrice,
		WaterPrice:    contract.Room.WaterPerUnit * water_unit,
		ElecPrice:     contract.Room.ElecPerUnit * elec_unit,
		WaterUnit:     water_unit,
		PrevWaterUnit: prev_water_unit,
		CurWaterUnit:  req.CurWaterUnit,

		ElecUnit:     elec_unit,
		PrevElecUnit: prev_elec_unit,
		CurElecUnit:  req.CurElecUnit,
		TotalAmount:  contract.Room.RentPrice + (contract.Room.WaterPerUnit * water_unit) + (contract.Room.ElecPerUnit * elec_unit),
	}
	return s.repo.Create(dto)
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
	if len(*contract.Invoice) > 0 {
		lastItem := (*contract.Invoice)[len(*contract.Invoice)-1]
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
	// ส่งไปที่ pdf adapters ส่วนประมวลผล gen pdf
}
