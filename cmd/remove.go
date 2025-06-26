package cmd

import (
	"context"
	"fmt"
	"github.com/ignite/cli/v29/ignite/services/plugin"
	inotify "ignite-notify/internal"
)

// Remove handles the 'notify rm' command
// Uses inotify.Subscription and helpers from internal/config.go
func Remove(ctx context.Context, c *plugin.ExecutedCommand) error {
	name := ""
	for _, f := range c.Flags {
		if f.Name == "name" {
			name = f.Value
		}
	}
	if name == "" && len(c.Args) > 0 {
		name = c.Args[0]
	}
	if name == "" {
		return fmt.Errorf("subscription name is required (use --name or as argument)")
	}

	file, err := inotify.GetConfigPath()
	if err != nil {
		return err
	}
	subs, err := inotify.LoadSubscriptions(file)
	if err != nil {
		return err
	}
	found := false
	newSubs := make([]inotify.Subscription, 0, len(subs))
	for _, s := range subs {
		if s.Name == name {
			found = true
			continue
		}
		newSubs = append(newSubs, s)
	}
	if !found {
		fmt.Printf("No subscription named '%s' found.\n", name)
		return nil
	}
	if err := inotify.SaveSubscriptions(file, newSubs); err != nil {
		return err
	}
	fmt.Printf("Subscription '%s' removed.\n", name)
	return nil
}
