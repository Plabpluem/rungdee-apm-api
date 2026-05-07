package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	usecases "rungdee-apm-api/internal/usecases/line"
	"rungdee-apm-api/internal/usecases/line/dto"
)

type HttpLineRepository struct{}

func NewHttpLineRepository() usecases.LineRepository {
	return &HttpLineRepository{}
}

func (h *HttpLineRepository) Sendmessage(data dto.LineMessageDto) error {
	token := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	if token == "" {
		return fmt.Errorf("LINE_CHANNEL_ACCESS_TOKEN is not set")
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("line push failed (status %d): %s", resp.StatusCode, string(respBody))
	}

	return nil
}

func (h *HttpLineRepository) SendFlexMessage(data dto.LineFlesMessageDto) error {
	token := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	if token == "" {
		return fmt.Errorf("LINE_CHANNEL_ACCESS_TOKEN is not set")
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("line push failed (status %d): %s", resp.StatusCode, string(respBody))
	}

	return nil

}
