package adapter

import (
	"fmt"
	usecases "rungdee-apm-api/internal/usecases/invoice"
	"rungdee-apm-api/internal/usecases/invoice/dto"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type HttpInvoiceHandlers struct {
	invoiceUseCase usecases.InvoiceUseCase
}

func NewHttpInvoiceHandler(usecase usecases.InvoiceUseCase) *HttpInvoiceHandlers {
	return &HttpInvoiceHandlers{invoiceUseCase: usecase}
}

func (h *HttpInvoiceHandlers) Create(c fiber.Ctx) error {
	var dto dto.CreateInvoiceDto

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}

	invoice, err := h.invoiceUseCase.Create(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create success", "data": invoice})
}

func (h *HttpInvoiceHandlers) Findall(c fiber.Ctx) error {
	query := new(dto.FilterInvoiceDto)
	if err := c.Bind().Query(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query"})
	}

	invoice, err := h.invoiceUseCase.Findall(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": invoice})

}

func (h *HttpInvoiceHandlers) Find(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	invoiceUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}

	invoice, err := h.invoiceUseCase.Find(&dto.FindInvoiceDto{UUid: invoiceUuid})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": invoice})

}

func (h *HttpInvoiceHandlers) Update(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	invoiceUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}
	var dto dto.UpdateInvoiceDto

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}
	dto.Uuid = invoiceUuid

	invoice, err := h.invoiceUseCase.Update(&dto)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update success", "data": invoice})

}

func (h *HttpInvoiceHandlers) GeneratePdf(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	invoiceUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}

	invoice, err := h.invoiceUseCase.Find(&dto.FindInvoiceDto{UUid: invoiceUuid})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	invoiceByte, err := h.invoiceUseCase.Generate(&dto.FindInvoiceDto{UUid: invoiceUuid})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	filename := fmt.Sprintf("invoice-%s-%s-%s.pdf", invoice.Contract.Room.Number, invoice.Contract.Customer.Name, time.Now().Format("dd-mm-YYYY"))
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", filename))
	return c.Send(invoiceByte)

}

func (h *HttpInvoiceHandlers) CreatePdf(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	invoiceUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}

	invoice_pdf, err := h.invoiceUseCase.CreatePdf(&dto.FindInvoiceDto{UUid: invoiceUuid})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": invoice_pdf})
}

func (h *HttpInvoiceHandlers) NotiToCustomer(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	invoiceUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}
	response, err := h.invoiceUseCase.NotiToCustomer(&dto.FindInvoiceDto{UUid: invoiceUuid})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": response})

}
