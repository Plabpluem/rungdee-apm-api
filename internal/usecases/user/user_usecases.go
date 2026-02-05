package usecases

import "rungdee-apm-api/internal/entities"

type UserUseCase interface {
	Create(user *entities.User) (*entities.User, error)
}

func NewUserService(repo UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

// สิ่งที่ส่งไป http
type UserService struct {
	repo UserRepository
}

func (s *UserService) Create(user *entities.User) (*entities.User, error) {
	return s.repo.Create(user)
}
