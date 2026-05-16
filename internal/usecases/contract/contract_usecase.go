package usecases

import (
	"rungdee-apm-api/internal/adapters/contract/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/contract/dto"
	storage_usecases "rungdee-apm-api/internal/usecases/storage"
)

type ContractUseCase interface {
	Create(dto *dto.CreateContractDto) (*entities.Contract, error)
	Findall(dto *dto.FilterContractDto) (*response.ContractPaginatedResponse, error)
	Findone(dto *dto.FindContractDto) (*entities.Contract, error)
	Update(dto *dto.UpdateContractDto) (*entities.Contract, error)
}

func NewContractService(repo ContractRepository, storageRepo storage_usecases.StorageRepository) ContractUseCase {
	return &ContractService{repo: repo, storageRepo: storageRepo}
}

type ContractService struct {
	repo        ContractRepository
	storageRepo storage_usecases.StorageRepository
}

func (s *ContractService) Create(dto *dto.CreateContractDto) (*entities.Contract, error) {
	return s.repo.Create(dto)
}

func (s *ContractService) Findall(dto *dto.FilterContractDto) (*response.ContractPaginatedResponse, error) {
	return s.repo.Findall(dto)
}

func (s *ContractService) Findone(dto *dto.FindContractDto) (*entities.Contract, error) {
	response, err := s.repo.Findone(dto)
	if err != nil {
		return nil, err
	}

	for index, item := range *response.Invoice {
		if item.LinkPdf != "" {
			image, err := s.storageRepo.GetUrl(item.LinkPdf, "pdf")
			if err != nil {
				return nil, err
			}
			(*response.Invoice)[index].LinkPdf = image
		}
	}
	return response, nil
}

func (s *ContractService) Update(dto *dto.UpdateContractDto) (*entities.Contract, error) {
	return s.repo.Update(dto)
}
