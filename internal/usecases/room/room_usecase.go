package usecases

import (
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/room/dto"
)

type RoomUseCase interface {
	Create(dto *dto.CreateRoomDto) (*entities.Room, error)
	FindAll() ([]*entities.Room, error)
	Update(dto *dto.UpdateRoomDto) (*entities.Room, error)
	Find(dto *dto.FindRoomDto) (*entities.Room, error)
}

func NewRoomService(repo RoomRepository) RoomUseCase {
	return &RoomService{repo: repo}
}

type RoomService struct {
	repo RoomRepository
}

func (s *RoomService) Create(dto *dto.CreateRoomDto) (*entities.Room, error) {
	return s.repo.Create(dto)
}

func (s *RoomService) FindAll() ([]*entities.Room, error) {
	return s.repo.FindAll()
}

func (s *RoomService) Find(dto *dto.FindRoomDto) (*entities.Room, error) {
	return s.repo.Find(dto)
}

func (s *RoomService) Update(dto *dto.UpdateRoomDto) (*entities.Room, error) {
	return s.repo.Update(dto)
}
