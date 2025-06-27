package sink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Sink interface {
	Send(msg string) error
}

type StdoutSink struct{}
func printFlat(prefix string, v interface{}) {
	switch val := v.(type) {
	case map[string]interface{}:
		for k, v2 := range val {
			printFlat(prefix+k+".", v2)
		}
	case []interface{}:
		for i, v2 := range val {
			printFlat(fmt.Sprintf("%s%d.", prefix, i), v2)
		}
	default:
		fmt.Printf("%s%v\n", prefix, val)
	}
}

func (s *StdoutSink) Send(msg string) error {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(msg), &data); err == nil {
		if result, ok := data["result"]; ok {
			printFlat("", result)
			return nil
		}
	}
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
func cleanPrefix(key string) string {
	if strings.HasPrefix(key, "data.value.") {
		return key[len("data.value."):]
	}
	if strings.HasPrefix(key, "data.") {
		return key[len("data."):]
	}
	return key
}

func flatAll(prefix string, v interface{}, out *map[string]string) {
	switch val := v.(type) {
	case map[string]interface{}:
		for k, v2 := range val {
			flatAll(prefix+k+".", v2, out)
		}
	case []interface{}:
		for i, v2 := range val {
			flatAll(fmt.Sprintf("%s%d.", prefix, i), v2, out)
		}
	default:
		key := prefix[:len(prefix)-1] // enlève le dernier point
		key = cleanPrefix(key)
		(*out)[key] = fmt.Sprintf("%v", val)
	}
}

func (s *TelegramSink) Send(msg string) error {
	var data map[string]interface{}
	text := msg
	if err := json.Unmarshal([]byte(msg), &data); err == nil {
		if result, ok := data["result"]; ok {
			out := map[string]string{}
			flatAll("", result, &out)
			lines := ""
			for k, v := range out {
				lines += fmt.Sprintf("%s: %s\n", k, v)
			}
			if lines != "" {
				text = lines
			}
		}
	}
	payload := map[string]interface{}{"text": text}
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
