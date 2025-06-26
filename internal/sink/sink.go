package sink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Sink interface {
	Send(msg string) error
}

type StdoutSink struct{}
func (s *StdoutSink) Send(msg string) error {
	fmt.Println(msg)
	return nil
}

type SlackSink struct {
	Webhook string
}
func (s *SlackSink) Send(msg string) error {
	payload := []byte(`{"text":` + fmt.Sprintf("%q", msg) + `}`)
	resp, err := http.Post(s.Webhook, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Slack webhook returned status %d", resp.StatusCode)
	}
	return nil
}

type TelegramSink struct {
	APIURL string
}
func (s *TelegramSink) Send(msg string) error {
	payload := map[string]interface{}{"text": msg}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(s.APIURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Telegram API returned status %d", resp.StatusCode)
	}
	return nil
}
