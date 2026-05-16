package dto

type LineMessageDto struct {
	To      string `json:"to"`
	Message any    `json:"message"`
}
