package usecases

import (
	"io"
	"rungdee-apm-api/internal/entities"
)

type StorageRepository interface {
	Save(file io.Reader, filename, contentType string) (*entities.StorageResponse, error)
	GetUrl(fileName string, subFolder string) (string, error)
}
