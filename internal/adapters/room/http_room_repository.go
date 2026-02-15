package adapters

import (
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/room"

	"github.com/gofiber/fiber/v3"
)

type HttpRoomHandler struct {
	roomUseCase usecases.RoomUseCase
}

func NewHttpRoomRepository(usecase usecases.RoomUseCase) *HttpRoomHandler {
	return &HttpRoomHandler{roomUseCase: usecase}
}

func (h *HttpRoomHandler) Create(c fiber.Ctx) error {
	var dto entities.Room
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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create", "data": room})
}
