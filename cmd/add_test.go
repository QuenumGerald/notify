package cmd

import (
	"context"
	"testing"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func TestAdd_Minimal(t *testing.T) {
	ctx := context.Background()
	c := &plugin.ExecutedCommand{
		Flags: []*plugin.FlagValue{
			{Name: "name", Value: "testadd"},
			{Name: "query", Value: "tm.event='NewBlock'"},
			{Name: "node", Value: "ws://localhost:26657"},
			{Name: "sink", Value: "stdout"},
		},
	}
	err := Add(ctx, c)
	if err != nil {
		t.Errorf("Add failed: %v", err)
	}
}
