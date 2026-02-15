package adapter

import (
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/user"
	"rungdee-apm-api/internal/usecases/user/dto"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type HttpUserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewHttpUserRepository(usecase usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: usecase}
}

func (h *HttpUserHandler) Signup(c fiber.Ctx) error {
	var dto entities.User

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	dto.Password = string(hashPassword)

	_, err = h.userUseCase.Create(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "create success", "statusCode": 201})
}

func (h *HttpUserHandler) Login(c fiber.Ctx) error {
	var dto dto.LoginDto
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid req"})
	}

	user, err := h.userUseCase.Login(&dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "password not match"})
	}

	claims := jwt.MapClaims{
		"username":  user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"user_uuid": user.Uuid,
		"role":      user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("supersecret"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"data": user, "token": t, "statusCode": 200})

}
