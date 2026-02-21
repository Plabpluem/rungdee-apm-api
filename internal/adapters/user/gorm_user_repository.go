package adapter

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/user"
	"rungdee-apm-api/internal/usecases/user/dto"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) usecases.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User

	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("username %s not found", username)
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Create(dto *dto.SignUpDto) (*entities.User, error) {
	user := &entities.User{
		Username: dto.Username,
		Password: dto.Password,
		Name:     dto.Name,
		Role:     dto.Role,
	}

	err := r.db.Create(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *GormUserRepository) FindAll() ([]*entities.User, error) {
	var user []*entities.User
	var total int64

	db := r.db.Model(user)
	db.Count(&total)

	err := db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *GormUserRepository) Find(dto *dto.FindUserDto) (*entities.User, error) {
	var user *entities.User

	db := r.db.Where("uuid = ?", dto.Uuid).Model(user)
	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *GormUserRepository) Update(dto *dto.UpdateUserDto) (*entities.User, error) {
	var user *entities.User

	db := r.db.Where("uuid = ?", dto.Uuid).Model(user)
	if err := db.Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil

}
