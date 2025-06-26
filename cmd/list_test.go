package cmd

import (
	"context"
	"testing"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func TestList_Minimal(t *testing.T) {
	ctx := context.Background()
	c := &plugin.ExecutedCommand{}
	err := List(ctx, c)
	if err != nil {
		t.Errorf("List failed: %v", err)
	}
}
