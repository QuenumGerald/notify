package subscriber

import (
	"context"
	"log"
	"strings"
	"time"
	"github.com/gorilla/websocket"
	"ignite-notify/internal/config"
	"ignite-notify/internal/sink"
)

type SubscriptionRunner struct {
	Sub  config.Subscription
	Sink sink.Sink
}

func (r *SubscriptionRunner) Run(ctx context.Context) {
	retry := 0
	for {
		if ctx.Err() != nil {
			return
		}
		wsURL := r.Sub.Node
		if strings.HasPrefix(wsURL, "tcp://") {
			wsURL = strings.Replace(wsURL, "tcp://", "ws://", 1)
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
				"query": r.Sub.Query,
			},
		}
		if err := c.WriteJSON(req); err != nil {
			log.Printf("[notify] WriteJSON error: %v\n", err)
			c.Close()
			time.Sleep(2 * time.Second)
			continue
		}
		log.Printf("[notify] Subscribed to %s: %s\n", r.Sub.Node, r.Sub.Query)
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
					break
				}
				r.Sink.Send(string(msg))
			}
		}
	}
}
