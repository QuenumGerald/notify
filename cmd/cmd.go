package cmd

import "github.com/ignite/cli/v29/ignite/services/plugin"

// GetCommands returns the list of ignite-notify app commands.
func GetCommands() []*plugin.Command {
	return []*plugin.Command{
		{
			Use:   "ignite-notify [command]",
			Short: "ignite-notify is an awesome Ignite application!",
			Commands: []*plugin.Command{
				{
					Use:   "hello",
					Short: "Say hello to the world of ignite!",
				},
			},
		},
	}
}
