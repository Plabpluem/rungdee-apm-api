package adapters

import (
	usecases "rungdee-apm-api/internal/usecases/storage"

	"github.com/gofiber/fiber/v3"
)

type HttpStorageHandler struct {
	storageUseCase usecases.StorageRepository
}

func NewHttpStorageHandler(usecases usecases.StorageRepository) *HttpStorageHandler {
	return &HttpStorageHandler{storageUseCase: usecases}
}

func (h *HttpStorageHandler) UploadFile(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	defer f.Close()

	url, err := h.storageUseCase.Save(f, file.Filename, file.Header.Get("Content-Type"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "upload success",
		"statusCode": 201,
		"url":        url,
	})
}
