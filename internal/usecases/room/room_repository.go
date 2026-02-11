package usecases

import "rungdee-apm-api/internal/entities"

type RoomRepository interface {
	Create(dto *entities.Room) (*entities.Room, error)
	FindAll() ([]*entities.Room, error)
}
