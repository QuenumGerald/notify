// SPDX-License-Identifier: MIT
package main

import (
	hplugin "github.com/hashicorp/go-plugin"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func main() {
	hplugin.Serve(&hplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig(),
		Plugins: map[string]hplugin.Plugin{
			"ignite-notify": plugin.NewGRPC(&notifyApp{}),
		},
		GRPCServer: hplugin.DefaultGRPCServer,
	})
}