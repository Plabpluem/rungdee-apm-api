package dto

import "github.com/google/uuid"

type FindRoomDto struct {
	Uuid uuid.UUID `json:"uuid"`
}
