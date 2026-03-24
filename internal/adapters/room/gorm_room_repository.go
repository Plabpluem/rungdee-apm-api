package adapters

import (
	"errors"
	"fmt"
	"rungdee-apm-api/internal/adapters/room/response"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/room"
	"rungdee-apm-api/internal/usecases/room/dto"
	"rungdee-apm-api/pkg"

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

func (r *GormRoomRepository) FindAll(dto *dto.FilterRoomDto) (*response.RoomPaginatedResponse, error) {
	var room []*entities.Room
	var total int64

	pagination := pkg.Pagination{
		Page:    dto.Page,
		PerPage: dto.PerPage,
	}

	db := r.db.Model(&entities.Room{})
	db.Count(&total)

	err := db.Limit(pagination.GetPerPage()).Offset(pagination.GetOffSet()).Find(&room).Error
	if err != nil {
		return nil, err
	}
	return &response.RoomPaginatedResponse{
		Data:       room,
		Total:      total,
		Page:       pagination.GetPage(),
		PerPage:    pagination.GetPerPage(),
		TotalPages: int((total + int64(pagination.GetPerPage()) - 1) / int64(pagination.GetPerPage())),
	}, nil
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

func (r *GormRoomRepository) FindAllDropdown() ([]*entities.Room, error) {
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
