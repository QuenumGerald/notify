package cmd

import (
	"context"
	"fmt"
	"github.com/ignite/cli/v29/ignite/services/plugin"
	"ignite-notify/internal/config"
)

// List handles the 'notify ls' command
// Uses config.Subscription and helpers from internal/config.go
func List(ctx context.Context, c *plugin.ExecutedCommand) error {
	file, err := config.GetConfigPath()
	if err != nil {
		return err
	}
	subs, err := config.LoadSubscriptions(file)
	if err != nil {
		return err
	}
	if len(subs) == 0 {
		fmt.Println("No subscriptions found.")
		return nil
	}
	fmt.Printf("%-16s %-26s %-22s %-8s %s\n", "NAME", "NODE", "QUERY", "SINK", "WEBHOOK")
	for _, s := range subs {
		fmt.Printf("%-16s %-26s %-22s %-8s %s\n", s.Name, s.Node, s.Query, s.Sink, s.Webhook)
	}
	return nil
}
