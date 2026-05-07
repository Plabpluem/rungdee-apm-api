package usecases

import "rungdee-apm-api/internal/usecases/line/dto"

type LineRepository interface {
	Sendmessage(dto dto.LineMessageDto) error
	SendFlexMessage(dto dto.LineFlesMessageDto) error
}
