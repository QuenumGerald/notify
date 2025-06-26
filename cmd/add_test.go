package cmd

import (
	"context"
	"testing"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func TestAdd_Minimal(t *testing.T) {
	ctx := context.Background()
	c := &plugin.ExecutedCommand{
		Flags: []*plugin.Flag{
			{Name: "name", Type: 1, Value: "testadd"},
			{Name: "query", Type: 1, Value: "tm.event='NewBlock'"},
			{Name: "node", Type: 1, Value: "ws://localhost:26657"},
			{Name: "sink", Type: 1, Value: "stdout"},
		},
	}
	err := Add(ctx, c)
	if err != nil {
		t.Errorf("Add failed: %v", err)
	}
}
