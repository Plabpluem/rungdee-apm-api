package usecases

import (
	"bytes"
	"fmt"
	"rungdee-apm-api/internal/adapters/invoice/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/invoice/dto"
	usecases "rungdee-apm-api/internal/usecases/line"
	dto_line "rungdee-apm-api/internal/usecases/line/dto"
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
	NotiToCustomer(req *dto.FindInvoiceDto) (*entities.Invoice, error)
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

	var elec_unit float64

	if req.CurElecUnit < prev_elec_unit {
		first_round := 9999 - prev_elec_unit
		last_round := 1 + req.CurElecUnit
		elec_unit = first_round + last_round
	} else {
		elec_unit = req.CurElecUnit - prev_elec_unit
	}

	var water_unit, water_price, cur_water_unit float64

	if req.IsUnlimitWater != nil && *req.IsUnlimitWater {
		water_unit = 0
		water_price = 200
		cur_water_unit = prev_water_unit
	} else {
		if req.CurWaterUnit == nil {
			return nil, fmt.Errorf("กรุณากรอก หน่วยน้ำปัจจุบัน")
		}
		cur_water_unit = *req.CurWaterUnit
		water_unit = cur_water_unit - prev_water_unit
		water_price = contract.Room.WaterPerUnit * water_unit
	}

	data_invoice := &entities.Invoice{
		ContractId:    contract.ID,
		RentPrice:     contract.Room.RentPrice,
		WaterPrice:    water_price,
		ElecPrice:     contract.Room.ElecPerUnit * elec_unit,
		ElecPerUnit:   contract.Room.ElecPerUnit,
		WaterPerUnit:  contract.Room.WaterPerUnit,
		WaterUnit:     &water_unit,
		PrevWaterUnit: prev_water_unit,
		CurWaterUnit:  cur_water_unit,

		ElecUnit:       elec_unit,
		PrevElecUnit:   prev_elec_unit,
		CurElecUnit:    req.CurElecUnit,
		TotalAmount:    contract.Room.RentPrice + (water_price) + (contract.Room.ElecPerUnit * elec_unit),
		IsUnlimitWater: req.IsUnlimitWater,
	}
	invoice, err := s.repo.Create(data_invoice)
	if err != nil {
		return nil, err
	}

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
	response, err := s.repo.Findall(dto)
	if err != nil {
		return nil, err
	}

	for index, item := range response.Data {
		if item.LinkPdf != "" {
			image, err := s.storageRepo.GetUrl(item.LinkPdf, "pdf")
			if err != nil {
				return nil, err
			}
			response.Data[index].LinkPdf = image
		}
	}
	return response, nil
}

func (s *InvoiceService) Find(dto *dto.FindInvoiceDto) (*entities.Invoice, error) {
	return s.repo.Find(dto)
}

func (s *InvoiceService) Update(req *dto.UpdateInvoiceDto) (*entities.Invoice, error) {
	contract, err := s.contractRepo.FindById(req.ContractId)
	if err != nil {
		return nil, err
	}

	prev_invoice, err := s.repo.Find(&dto.FindInvoiceDto{UUid: req.Uuid})
	if err != nil {
		return nil, err
	}

	var prev_elec_unit float64
	var prev_water_unit float64

	if len(*contract.Invoice) > 1 {
		var index_previous int
		for index, item := range *contract.Invoice {
			if item.ID == prev_invoice.ID {
				index_previous = index - 1
			}
		}
		if index_previous < 0 {
			prev_elec_unit = contract.StartElecUnit
			prev_water_unit = contract.StartWaterUnit
		} else {
			lastItem := (*contract.Invoice)[index_previous]
			prev_elec_unit = lastItem.CurElecUnit
			prev_water_unit = lastItem.CurWaterUnit
		}
	} else {
		prev_elec_unit = req.PrevElecUnit
		prev_water_unit = req.PrevWaterUnit
	}

	var elec_unit float64

	if req.CurElecUnit < prev_elec_unit {
		first_round := 9999 - prev_elec_unit
		last_round := 1 + req.CurElecUnit
		elec_unit = first_round + last_round
	} else {
		elec_unit = req.CurElecUnit - prev_elec_unit
	}

	var water_unit, water_price, cur_water_unit float64

	if req.IsUnlimitWater != nil && *req.IsUnlimitWater {
		water_unit = 0
		water_price = 200
		cur_water_unit = prev_water_unit
	} else {
		if req.CurWaterUnit == nil {
			return nil, fmt.Errorf("กรุณากรอก หน่วยน้ำปัจจุบัน")
		}
		cur_water_unit = *req.CurWaterUnit
		water_unit = cur_water_unit - prev_water_unit
		water_price = contract.Room.WaterPerUnit * water_unit
	}

	dto_update := &entities.Invoice{
		Uuid:          req.Uuid,
		ContractId:    contract.ID,
		RentPrice:     contract.Room.RentPrice,
		WaterPrice:    water_price,
		ElecPrice:     contract.Room.ElecPerUnit * elec_unit,
		ElecPerUnit:   contract.Room.ElecPerUnit,
		WaterPerUnit:  contract.Room.WaterPerUnit,
		WaterUnit:     &water_unit,
		PrevWaterUnit: prev_water_unit,
		CurWaterUnit:  cur_water_unit,

		ElecUnit:       elec_unit,
		PrevElecUnit:   prev_elec_unit,
		CurElecUnit:    req.CurElecUnit,
		TotalAmount:    contract.Room.RentPrice + water_price + (contract.Room.ElecPerUnit * elec_unit),
		IsUnlimitWater: req.IsUnlimitWater,
	}

	update, err := s.repo.Update(dto_update)
	if err != nil {
		return nil, err
	}

	var check_equal_water bool
	if req.IsUnlimitWater != nil && *req.IsUnlimitWater {
		check_equal_water = true
	} else if cur_water_unit == prev_invoice.CurWaterUnit {
		check_equal_water = true
	} else {
		check_equal_water = false
	}

	if req.CurElecUnit == prev_invoice.CurElecUnit && check_equal_water {
		return update, nil
	}

	invoice_dto := &dto.FindInvoiceDto{
		UUid: update.Uuid,
	}

	pdf_link, err := s.CreatePdf(invoice_dto)
	if err != nil {
		return nil, err
	}

	dto_update = &entities.Invoice{
		Uuid:    update.Uuid,
		LinkPdf: pdf_link.BaseUrl,
	}
	return s.repo.Update(dto_update)

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

	uploadPdf, err := s.storageRepo.Save(bytes.NewReader(generate_pdf), fmt.Sprintf("invoice-%s-%s-%s.pdf", invoice.Contract.Room.Number, invoice.Contract.Customer.Name, time.Now().Format("02-01-2006")), "application/pdf")
	if err != nil {
		return nil, err
	}

	return uploadPdf, nil
}

func (s *InvoiceService) NotiToCustomer(req *dto.FindInvoiceDto) (*entities.Invoice, error) {
	invoice, err := s.repo.Find(req)
	if err != nil {
		return nil, err
	}

	if invoice.Contract.Customer.LineUserId == "" {
		return nil, fmt.Errorf("ไม่พบเลข lineId")
	}

	pdf_link, err := s.storageRepo.GetUrl(invoice.LinkPdf, "pdf")
	if err != nil {
		return nil, err
	}

	data := &dto_line.LineFlesMessageDto{
		To: invoice.Contract.Customer.LineUserId,
		Messages: []map[string]any{
			{
				"type":    "flex",
				"altText": "ใบแจ้งหนี้",
				"contents": map[string]any{
					"type": "bubble",
					"header": map[string]any{
						"type":   "box",
						"layout": "vertical",
						"contents": []map[string]any{
							{
								"type":   "text",
								"text":   "ใบแจ้งหนี้",
								"weight": "bold",
								"size":   "xl",
								"align":  "center",
							},
						},
					},
					"body": map[string]any{
						"type":   "box",
						"layout": "vertical",
						"contents": []map[string]any{
							{
								"type": "text",
								"text": "คุณสามารถดูลายละเอียดได้ที่ปุ่มด้านล่างนี้",
								"wrap": true,
							},
						},
					},
					"footer": map[string]any{
						"type":   "box",
						"layout": "vertical",
						"contents": []map[string]any{
							{
								"type":  "button",
								"style": "primary",
								"action": map[string]any{
									"type":  "uri",
									"label": "ดูใบแจ้งหนี้",
									"uri":   pdf_link,
								},
							},
						},
					},
				},
			},
		},
	}
	if err := s.lineRepo.SendFlexMessage(*data); err != nil {
		return nil, err
	}
	return invoice, nil
}
