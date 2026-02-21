package usecases

import (
	"fmt"
	"os"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/user/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Create(user *dto.SignUpDto) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Find(dto *dto.FindUserDto) (*entities.User, error)
	Update(dto *dto.UpdateUserDto) (*entities.User, error)
	Login(dto *dto.LoginDto) (string, error)
}

func NewUserService(repo UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) Create(user *dto.SignUpDto) (*entities.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashPassword)

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

func (s *UserService) Login(loginDto *dto.LoginDto) (string, error) {
	user, err := s.repo.FindByUsername(loginDto.Username)
	if err != nil {
		return "", fmt.Errorf("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return "", fmt.Errorf("password not match")
	}

	claims := jwt.MapClaims{
		"username":  user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"user_uuid": user.Uuid,
		"role":      user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}
