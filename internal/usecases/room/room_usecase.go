package usecases

import (
	"rungdee-apm-api/internal/adapters/room/response"
	"rungdee-apm-api/internal/entities"
	"rungdee-apm-api/internal/usecases/room/dto"
	responsePkg "rungdee-apm-api/pkg/response"
)

type RoomUseCase interface {
	Create(dto *dto.CreateRoomDto) (*entities.Room, error)
	FindAll(dto *dto.FilterRoomDto) (*response.RoomPaginatedResponse, error)
	Update(dto *dto.UpdateRoomDto) (*entities.Room, error)
	Find(dto *dto.FindRoomDto) (*entities.Room, error)
	FindalllDropdown() ([]*responsePkg.FindAllDropdownResponse, error)
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

func (s *RoomService) FindAll(dto *dto.FilterRoomDto) (*response.RoomPaginatedResponse, error) {
	return s.repo.FindAll(dto)
}

func (s *RoomService) Find(dto *dto.FindRoomDto) (*entities.Room, error) {
	return s.repo.Find(dto)
}

func (s *RoomService) Update(dto *dto.UpdateRoomDto) (*entities.Room, error) {
	return s.repo.Update(dto)
}

func (s *RoomService) FindalllDropdown() ([]*responsePkg.FindAllDropdownResponse, error) {
	room, err := s.repo.FindAllDropdown()
	if err != nil {
		return nil, err
	}

	room_response := make([]*responsePkg.FindAllDropdownResponse, 0, len(room))

	for _, item := range room {
		room_response = append(room_response, &responsePkg.FindAllDropdownResponse{
			Label: item.Number,
			Value: item.ID,
		})
	}
	return room_response, err
}
