package adapter

import (
	usecases "rungdee-apm-api/internal/usecases/user"
	"rungdee-apm-api/internal/usecases/user/dto"

	"github.com/gofiber/fiber/v3"
)

type HttpUserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewHttpUserHandler(usecase usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: usecase}
}

func (h *HttpUserHandler) Signup(c fiber.Ctx) error {
	var user dto.SignUpDto

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}

	_, err := h.userUseCase.Create(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create success", "statusCode": 201})
}

func (h *HttpUserHandler) Login(c fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := c.Bind().Body(&loginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}

	token, err := h.userUseCase.Login(&loginDto)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "login success", "token": token})
}
