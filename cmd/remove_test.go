package cmd

import (
	"context"
	"testing"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func TestRemove_Minimal(t *testing.T) {
	ctx := context.Background()
	c := &plugin.ExecutedCommand{
		Flags: []*plugin.FlagValue{{Name: "name", Value: "testadd"}},
	}
	err := Remove(ctx, c)
	if err != nil {
		t.Errorf("Remove failed: %v", err)
	}
}
