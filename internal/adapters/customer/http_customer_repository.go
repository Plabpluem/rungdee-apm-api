package adapter

import (
	usecases "rungdee-apm-api/internal/usecases/customer"
	"rungdee-apm-api/internal/usecases/customer/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

func (h *HttpCustomerHandler) Findall(c fiber.Ctx) error {

	query := new(dto.FilterCustomerDto)

	if err := c.Bind().Query(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query"})
	}

	customer, err := h.customerUseCase.Findall(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": customer})
}

func (h *HttpCustomerHandler) Find(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	customerUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}

	customer, err := h.customerUseCase.Find(&dto.FindCustomerDto{UUid: customerUuid})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": customer})
}

func (h HttpCustomerHandler) Update(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	customerUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}

	var dto dto.UpdateCustomerDto

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}
	dto.Uuid = customerUuid

	customer, err := h.customerUseCase.Update(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update success", "data": customer})
}
