package dto

type CreateRoomDto struct {
	Number       string  `json:"number"`
	RentPrice    float64 `json:"rent_price"`
	WaterPerUnit float64 `json:"water_per_unit"`
	ElecPerUnit  float64 `json:"elec_per_unit"`
}
