package cmd

import (
	"context"
	"fmt"
	"ignite-notify/internal/config"
	"ignite-notify/internal/runner"

	"github.com/ignite/cli/v29/ignite/services/plugin"
)

// Run handles the 'notify run' command
func Run(ctx context.Context, c *plugin.ExecutedCommand) error {
	file, err := config.GetConfigPath()
	if err != nil {
		return err
	}
	subs, err := config.LoadSubscriptions(file)
	if err != nil {
		return err
	}
	if len(subs) == 0 {
		fmt.Println("No subscriptions found. Use 'notify add' to create one.")
		return nil
	}
	r := runner.Runner{Subs: subs}
	return r.Start(ctx)
}

// AutoRun lance le runner notify en mode automatique (hook)
func AutoRun(ctx context.Context) error {
	return Run(ctx, nil)
}
