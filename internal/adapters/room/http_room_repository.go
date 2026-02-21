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
	room, err := h.roomUseCase.FindAll()
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
