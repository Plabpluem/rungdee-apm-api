package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/contract/dto"
)

type ContractUseCase interface {
	Create(dto *dto.CreateContractDto) (*entities.Contract, error)
	Findall() ([]*entities.Contract, error)
	Findone(dto *dto.FindContractDto) (*entities.Contract, error)
	Update(dto *dto.UpdateContractDto) (*entities.Contract, error)
}

func NewContractService(repo ContractRepository) ContractUseCase {
	return &ContractService{repo: repo}
}

type ContractService struct {
	repo ContractRepository
}

func (s *ContractService) Create(dto *dto.CreateContractDto) (*entities.Contract, error) {
	return s.repo.Create(dto)
}

func (s *ContractService) Findall() ([]*entities.Contract, error) {
	return s.repo.Findall()
}

func (s *ContractService) Findone(dto *dto.FindContractDto) (*entities.Contract, error) {
	return s.repo.Findone(dto)
}

func (s *ContractService) Update(dto *dto.UpdateContractDto) (*entities.Contract, error) {
	return s.repo.Update(dto)
}
