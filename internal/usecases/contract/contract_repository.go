package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/contract/dto"
)

type ContractRepository interface {
	Create(dto *dto.CreateContractDto) (*entities.Contract, error)
	Findall() ([]*entities.Contract, error)
	Findone(dto *dto.FindContractDto) (*entities.Contract, error)
	Update(dto *dto.UpdateContractDto) (*entities.Contract, error)
}
