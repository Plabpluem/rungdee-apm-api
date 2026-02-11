package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/user/dto"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Find(dto *dto.FindUserDto) (*entities.User, error)
	Update(dto *dto.UpdateUserDto) (*entities.User, error)
}
