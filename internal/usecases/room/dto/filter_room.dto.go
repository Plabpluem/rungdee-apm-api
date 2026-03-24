package dto

type FilterRoomDto struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Query   string `json:"query"`
}
