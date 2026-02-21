package adapters

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/room"
	"rungdee-apm-api/internal/usecases/room/dto"

	"gorm.io/gorm"
)

type GormRoomRepository struct {
	db *gorm.DB
}

func NewGormRoomRepository(db *gorm.DB) usecases.RoomRepository {
	return &GormRoomRepository{db: db}
}

func (r *GormRoomRepository) Create(dto *dto.CreateRoomDto) (*entities.Room, error) {
	room := &entities.Room{
		Number:       dto.Number,
		RentPrice:    dto.RentPrice,
		WaterPerUnit: dto.WaterPerUnit,
		ElecPerUnit:  dto.ElecPerUnit,
	}
	err := r.db.Create(room).Error

	if err != nil {
		return nil, err
	}
	return room, err
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

func (r *GormRoomRepository) Find(dto *dto.FindRoomDto) (*entities.Room, error) {
	var room entities.Room

	err := r.db.Where("uuid = ?", dto.Uuid).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("room not found")
		}
		return nil, err
	}
	return &room, nil
}

func (r *GormRoomRepository) Update(dto *dto.UpdateRoomDto) (*entities.Room, error) {
	var room entities.Room
	err := r.db.Where("uuid = ?", dto.Uuid).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("room not found")
		}
		return nil, err
	}

	err = r.db.Model(&room).Updates(entities.Room{
		Number:       dto.Number,
		RentPrice:    dto.RentPrice,
		WaterPerUnit: dto.WaterPerUnit,
		ElecPerUnit:  dto.ElecPerUnit,
	}).Error

	if err != nil {
		return nil, err
	}

	return &room, err
}
