package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/room/dto"
)

type RoomRepository interface {
	Create(dto *dto.CreateRoomDto) (*entities.Room, error)
	FindAll() ([]*entities.Room, error)
	Find(dto *dto.FindRoomDto) (*entities.Room, error)
	Update(dto *dto.UpdateRoomDto) (*entities.Room, error)
}
