package sink

import (
	"testing"
)

func TestStdoutSink_Send(t *testing.T) {
	s := &StdoutSink{}
	err := s.Send("hello world")
	if err != nil {
		t.Errorf("StdoutSink.Send failed: %v", err)
	}
}

func TestSlackSink_Send(t *testing.T) {
	s := &SlackSink{Webhook: "https://hooks.slack.com/services/invalid"}
	err := s.Send("test slack")
	if err == nil {
		t.Error("SlackSink.Send should fail with invalid webhook")
	}
}

func TestTelegramSink_Send(t *testing.T) {
	s := &TelegramSink{APIURL: "https://api.telegram.org/botINVALID/sendMessage?chat_id=123"}
	err := s.Send("test telegram")
	if err == nil {
		t.Error("TelegramSink.Send should fail with invalid API URL")
	}
}
