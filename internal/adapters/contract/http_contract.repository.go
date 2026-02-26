package adapters

import (
	usecases "rungdee-apm-api/internal/usecases/contract"
	"rungdee-apm-api/internal/usecases/contract/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type HttpContractHandler struct {
	contractUseCase usecases.ContractUseCase
}

func NewHttpContractHandler(usecase usecases.ContractUseCase) *HttpContractHandler {
	return &HttpContractHandler{contractUseCase: usecase}
}

func (h *HttpContractHandler) Findall(c fiber.Ctx) error {
	query := new(dto.FilterContractDto)

	if err := c.Bind().Query(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query"})
	}

	contract, err := h.contractUseCase.Findall(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": contract})

}

func (h *HttpContractHandler) Create(c fiber.Ctx) error {
	var dto dto.CreateContractDto

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})

	}

	contract, err := h.contractUseCase.Create(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create success", "data": contract})
}

func (h *HttpContractHandler) Find(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	contractUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}

	contract, err := h.contractUseCase.Findone(&dto.FindContractDto{Uuid: contractUuid})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": contract})

}

func (h *HttpContractHandler) Update(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	contractUuid, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}
	var dto dto.UpdateContractDto
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})

	}
	dto.Uuid = contractUuid

	contract, err := h.contractUseCase.Update(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update success", "data": contract})
}
