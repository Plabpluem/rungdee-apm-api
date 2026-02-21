package adapter

import (
	usecases "rungdee-apm-api/internal/usecases/customer"
	"rungdee-apm-api/internal/usecases/customer/dto"

	"github.com/gofiber/fiber/v3"
)

type HttpCustomerHandler struct {
	customerUseCase usecases.CustomerUseCase
}

func NewHttpCustomerHandler(usecase usecases.CustomerUseCase) *HttpCustomerHandler {
	return &HttpCustomerHandler{customerUseCase: usecase}
}

func (h *HttpCustomerHandler) Create(c fiber.Ctx) error {
	var dto dto.CreateCustomerDto

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})

	}

	customer, err := h.customerUseCase.Create(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create success", "data": customer})
}
