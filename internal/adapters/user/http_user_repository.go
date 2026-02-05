package adapter

import (
	usecases "rungdee-apm-api/internal/usecases/user"

	"github.com/gofiber/fiber/v3"
)

type HttpUserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewHttpUserRepository(usecase usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: usecase}
}

func (h *HttpUserHandler) CreateUser(c fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create"})
}
