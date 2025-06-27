// SPDX-License-Identifier: MIT
package main

import (
	"context"
	"fmt"

	"github.com/ignite/cli/v29/ignite/services/plugin"

	"ignite-notify/cmd" // vos helpers par commande
)

type notifyApp struct{}

/* ---------- Manifest : seulement la description ---------- */

func (notifyApp) Manifest(_ context.Context) (*plugin.Manifest, error) {
	return &plugin.Manifest{
		Name:       "notify",
		SharedHost: true,

		Commands: []*plugin.Command{
			{
				Use:   "add",
				Short: "Add a new subscription",
				Flags: []*plugin.Flag{
					{Name: "name", Type: plugin.FlagTypeString, Usage: "subscription name", Shorthand: "n"},
					{Name: "node", Type: plugin.FlagTypeString, Usage: "Tendermint RPC address", Shorthand: "N"}, // Set default in logic
					{Name: "query", Type: plugin.FlagTypeString, Usage: "event query"},
					{Name: "sink", Type: plugin.FlagTypeString, Usage: "stdout|slack"}, // Set default in logic
					{Name: "webhook", Type: plugin.FlagTypeString, Usage: "Slack webhook URL"},
				},
			},
			{Use: "run", Short: "Start all subscriptions"},
			{Use: "ls",  Short: "List subscriptions"},
			{Use: "rm [name]", Short: "Remove a subscription"},
		},
	}, nil
}

/* ---------- Command dispatcher ---------- */

func (notifyApp) Execute(ctx context.Context, c *plugin.ExecutedCommand, _ plugin.ClientAPI) error {
	switch c.Path {
	case "add", "ignite add":
		return cmd.Add(ctx, c)
	case "run", "ignite run":
		return cmd.Run(ctx, c)
	case "ls", "ignite ls":
		return cmd.List(ctx, c)
	case "rm", "ignite rm":
		return cmd.Remove(ctx, c)
	default:
		return fmt.Errorf("unknown command path: %s", c.Path)
	}
}

/* ---------- Hooks ---------- */

func (notifyApp) ExecuteHookPre(context.Context, *plugin.ExecutedHook, plugin.ClientAPI) error     { return nil }

func (notifyApp) ExecuteHookPost(ctx context.Context, h *plugin.ExecutedHook, _ plugin.ClientAPI) error {
	if h.Hook.GetName() == "auto-run" {
		return cmd.AutoRun(ctx)  // démarre le runner en arrière-plan
	}
	return nil
}

func (notifyApp) ExecuteHookCleanUp(context.Context, *plugin.ExecutedHook, plugin.ClientAPI) error { return nil }