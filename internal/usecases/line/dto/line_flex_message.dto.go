package dto

type LineFlesMessageDto struct {
	To       string           `json:"to"`
	Messages []map[string]any `json:"messages"`
}
