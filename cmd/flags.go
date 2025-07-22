package cmd

import (
	"github.com/ignite/cli/v29/ignite/services/plugin"
	"strings"
)

// flagValue returns the value of the first matching flag from the executed command.
// It compares flag names case-insensitively and also checks shorthand names.
func flagValue(c *plugin.ExecutedCommand, names ...string) string {
	for _, f := range c.Flags {
		name := strings.ToLower(strings.TrimLeft(f.Name, "-"))
		for _, n := range names {
			if name == strings.ToLower(n) {
				return f.Value
			}
		}
		if f.Shorthand != "" {
			for _, n := range names {
				if f.Shorthand == n {
					return f.Value
				}
			}
		}
	}
	return ""
}
