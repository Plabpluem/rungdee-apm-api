package dto

import (
	"rungdee-apm-api/internal/entities"
	"time"

	"github.com/google/uuid"
)

type UpdateContractDto struct {
	Uuid           uuid.UUID               `json:"uuid"`
	RoomId         uint                    `json:"room_id"`
	CustomerId     uint                    `gorm:"uniqueIndex" json:"customer_id"`
	StartDate      time.Time               `gorm:"type:date" json:"start_date"`
	EndDate        time.Time               `gorm:"type:date" json:"end_date"`
	Status         entities.StatusContract `gorm:"type:varchar(20);default:'active'" json:"status"`
	StartElecUnit  float64                 `json:"start_elec_unit"`
	StartWaterUnit float64                 `json:"start_water_unit"`
}
