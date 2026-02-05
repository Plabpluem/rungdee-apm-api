package usecases

import "rungdee-apm-api/internal/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
}
