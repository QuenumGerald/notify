package cmd

import (
	"context"
	"fmt"
	"github.com/ignite/cli/v29/ignite/services/plugin"
	"notify/internal/config"
)

// Add handles the 'notify add' command
// Uses config.Subscription and helpers from internal/config.go
func Add(ctx context.Context, c *plugin.ExecutedCommand) error {
	name, node, query, sink, webhook := "", "", "", "", ""
	for _, f := range c.Flags {
		switch f.Name {
		case "name":
			name = f.Value
		case "node":
			node = f.Value
		case "query":
			query = f.Value
		case "sink":
			sink = f.Value
		case "webhook":
			webhook = f.Value
		}
	}

	// Set defaults
	if node == "" {
		node = "tcp://localhost:26657"
	}
	if sink == "" {
		sink = "stdout"
	}
	if name == "" || query == "" {
		return fmt.Errorf("name and query are required")
	}

	sub := config.Subscription{
		Name:    name,
		Node:    node,
		Query:   query,
		Sink:    sink,
		Webhook: webhook,
	}

	file, err := config.GetConfigPath()
	if err != nil {
		return err
	}

	subs, err := config.LoadSubscriptions(file)
	if err != nil {
		return err
	}

	// Check for duplicate name
	for _, s := range subs {
		if s.Name == name {
			return fmt.Errorf("subscription with name '%s' already exists", name)
		}
	}

	subs = append(subs, sub)
	if err := config.SaveSubscriptions(file, subs); err != nil {
		return err
	}

	fmt.Printf("Added subscription '%s' (query: %s, sink: %s)\n", name, query, sink)
	return nil
}
