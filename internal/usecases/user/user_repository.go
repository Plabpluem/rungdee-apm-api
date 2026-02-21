package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/user/dto"
)

type UserRepository interface {
	Create(user *dto.SignUpDto) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Find(dto *dto.FindUserDto) (*entities.User, error)
	Update(dto *dto.UpdateUserDto) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
}
