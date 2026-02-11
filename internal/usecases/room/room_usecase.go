package usecases

import "rungdee-apm-api/internal/entities"

type RoomUseCase interface {
	Create(dto *entities.Room) (*entities.Room, error)
	FindAll() ([]*entities.Room, error)
}

func NewRoomService(repo RoomRepository) RoomUseCase {
	return &RoomService{repo: repo}
}

type RoomService struct {
	repo RoomRepository
}

func (s *RoomService) Create(dto *entities.Room) (*entities.Room, error) {
	return s.repo.Create(dto)
}

func (s *RoomService) FindAll() ([]*entities.Room, error) {
	return s.repo.FindAll()
}
