package adapters

import (
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/room"

	"gorm.io/gorm"
)

type GormRoomRepository struct {
	db *gorm.DB
}

func NewGormRoomRepository(db *gorm.DB) usecases.RoomRepository {
	return &GormRoomRepository{db: db}
}

func (r *GormRoomRepository) Create(dto *entities.Room) (*entities.Room, error) {
	err := r.db.Create(dto).Error

	if err != nil {
		return nil, err
	}
	return dto, err
}

func (r *GormRoomRepository) FindAll() ([]*entities.Room, error) {
	var room []*entities.Room
	var total int64

	db := r.db.Model(room)
	db.Count(&total)

	err := db.Find(&room).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}
