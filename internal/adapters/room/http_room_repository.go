package adapters

import (
	usecases "rungdee-apm-api/internal/usecases/room"
	"rungdee-apm-api/internal/usecases/room/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type HttpRoomHandler struct {
	roomUseCase usecases.RoomUseCase
}

func NewHttpRoomHandler(usecase usecases.RoomUseCase) *HttpRoomHandler {
	return &HttpRoomHandler{roomUseCase: usecase}
}

func (h *HttpRoomHandler) Create(c fiber.Ctx) error {
	var dto dto.CreateRoomDto
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}
	room, err := h.roomUseCase.Create(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create", "data": room})
}

func (h *HttpRoomHandler) FindAll(c fiber.Ctx) error {
	query := new(dto.FilterRoomDto)
	if err := c.Bind().Query(query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query"})
	}
	room, err := h.roomUseCase.FindAll(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": room})
}

func (h *HttpRoomHandler) Find(c fiber.Ctx) error {
	uuidParam := c.Params("id")
	roomUuid, err := uuid.Parse(uuidParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}
	room, err := h.roomUseCase.Find(&dto.FindRoomDto{Uuid: roomUuid})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": room})
}

func (h *HttpRoomHandler) Update(c fiber.Ctx) error {
	uuidParams := c.Params("id")
	roomParams, err := uuid.Parse(uuidParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid uuid error"})
	}
	var dto dto.UpdateRoomDto
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}
	dto.Uuid = roomParams
	room, err := h.roomUseCase.Update(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update success", "data": room})
}

func (h *HttpRoomHandler) FindallDropdown(c fiber.Ctx) error {
	room, err := h.roomUseCase.FindalllDropdown()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "fetch success", "data": room})

}
