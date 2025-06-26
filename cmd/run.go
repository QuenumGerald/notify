package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/gorilla/websocket"
	"github.com/ignite/cli/v29/ignite/services/plugin"
	inotify "ignite-notify/internal"
)

// Run handles the 'notify run' command
func Run(ctx context.Context, c *plugin.ExecutedCommand) error {
	file, err := inotify.GetConfigPath()
	if err != nil {
		return err
	}
	subs, err := inotify.LoadSubscriptions(file)
	if err != nil {
		return err
	}
	if len(subs) == 0 {
		fmt.Println("No subscriptions to run.")
		return nil
	}
	for _, s := range subs {
		fmt.Printf("[notify] Would start listening: name=%s node=%s query=%s sink=%s webhook=%s\n", s.Name, s.Node, s.Query, s.Sink, s.Webhook)
		go startSubscriptionWS(ctx, s)
	}
	return nil
}

func startSubscriptionWS(ctx context.Context, sub inotify.Subscription) {
	retry := 0
	for {
		if ctx.Err() != nil {
			return
		}
		var wsURL string
		if strings.HasPrefix(sub.Node, "tcp://") {
			wsURL = strings.Replace(sub.Node, "tcp://", "ws://", 1)
		} else {
			wsURL = sub.Node
		}
		if !strings.HasSuffix(wsURL, "/websocket") {
			wsURL = wsURL + "/websocket"
		}

		c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			log.Printf("[notify] WebSocket dial error: %v\n", err)
			retry++
			if retry >= 5 {
				log.Printf("[notify] Too many failures, waiting 10s before retrying...")
				time.Sleep(10 * time.Second)
				retry = 0
			} else {
				time.Sleep(2 * time.Second)
			}
			continue
		}
		retry = 0
		defer c.Close()

		req := map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "subscribe",
			"id":      "1",
			"params": map[string]interface{}{
				"query": sub.Query,
			},
		}
		if err := c.WriteJSON(req); err != nil {
			log.Printf("[notify] WriteJSON error: %v\n", err)
			c.Close()
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Printf("[notify] Subscribed to %s: %s\n", sub.Node, sub.Query)

		for {
			select {
			case <-ctx.Done():
				c.Close()
				return
			default:
				_, msg, err := c.ReadMessage()
				if err != nil {
					log.Printf("[notify] Read error: %v\n", err)
					c.Close()
					time.Sleep(2 * time.Second)
					break // sort du for interne, retente la reconnexion
				}

				// Sink dispatch
				switch sub.Sink {
				case "stdout":
					fmt.Printf("[notify][%s] %s\n", sub.Name, string(msg))
				case "slack":
					if sub.Webhook != "" {
						err := sendSlackNotification(sub.Webhook, string(msg))
						if err != nil {
							log.Printf("[notify][slack] error: %v\n", err)
						}
					}
				case "telegram":
					if sub.Webhook != "" {
						err := sendTelegramNotification(sub.Webhook, string(msg))
						if err != nil {
							log.Printf("[notify][telegram] error: %v\n", err)
						}
					}
				default:
					fmt.Printf("[notify] Sink '%s' not implemented yet.\n", sub.Sink)
				}
			}
		}
	}
}

// sendTelegramNotification sends a message to a Telegram bot API endpoint
// sub.Webhook doit être de la forme "https://api.telegram.org/bot<TOKEN>/sendMessage?chat_id=<ID>"
func sendTelegramNotification(apiURL, msg string) error {
	payload := map[string]interface{}{"text": msg}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Telegram API returned status %d", resp.StatusCode)
	}
	return nil
}

// sendSlackNotification sends a message to a Slack webhook URL
func sendSlackNotification(webhookURL, msg string) error {
	payload := []byte(`{"text":` + fmt.Sprintf("%q", msg) + `}`)
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Slack webhook returned status %d", resp.StatusCode)
	}
	return nil
}
