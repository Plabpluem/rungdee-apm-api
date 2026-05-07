package usecases

import (
	"io"
	"rungdee-apm-api/internal/entities"
)

type StorageUseCase interface {
	Save(file io.Reader, filename, contentType string) (*entities.StorageResponse, error)
	GetUrl(fileName string, subFolder string) (string, error)
}

type StorageService struct {
	repo StorageRepository
}

func NewStorageService(repo StorageRepository) StorageUseCase {
	return &StorageService{repo: repo}
}

func (s *StorageService) Save(file io.Reader, filename, contentType string) (*entities.StorageResponse, error) {
	return s.repo.Save(file, filename, contentType)
}

func (s *StorageService) GetUrl(fileName string, subFolder string) (string, error) {
	return s.repo.GetUrl(fileName, subFolder)
}
