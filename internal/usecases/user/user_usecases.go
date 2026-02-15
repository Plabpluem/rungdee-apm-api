package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/user/dto"
)

type UserUseCase interface {
	Create(user *entities.User) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Find(dto *dto.FindUserDto) (*entities.User, error)
	Update(dto *dto.UpdateUserDto) (*entities.User, error)
	Login(dto *dto.LoginDto) (*entities.User, error)
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
func (s *UserService) FindAll() ([]*entities.User, error) {
	return s.repo.FindAll()
}
func (s *UserService) Find(dto *dto.FindUserDto) (*entities.User, error) {
	return s.repo.Find(dto)
}
func (s *UserService) Update(dto *dto.UpdateUserDto) (*entities.User, error) {
	return s.repo.Update(dto)
}
func (s *UserService) Login(dto *dto.LoginDto) (*entities.User, error) {
	return s.repo.Login(dto)
}
