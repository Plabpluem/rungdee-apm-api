package usecases

import (
	"rungdee-apm-api/internal/adapters/invoice/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/invoice/dto"

	"github.com/google/uuid"
)

type InvoiceRepository interface {
	Create(dto *entities.Invoice) (*entities.Invoice, error)
	Find(dto *dto.FindInvoiceDto) (*entities.Invoice, error)
	Findall(dto *dto.FilterInvoiceDto) (*response.InvoicePaginatedResponse, error)
	Update(dto *entities.Invoice) (*entities.Invoice, error)
}

type ContractReader interface {
	FindByUuid(uuid uuid.UUID) (*entities.Contract, error)
	FindById(id uint) (*entities.Contract, error)
}

type PdfGenerate interface {
	Generate(dto *PdfData) ([]byte, error)
}

type PdfData struct {
	Invoice  *entities.Invoice
	Room     *entities.Room
	Customer *entities.Customer
}
