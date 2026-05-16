package dto

type LineMessageDto struct {
	To       string        `json:"to"`
	Messages []LineMessage `json:"messages"`
}

type LineMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
