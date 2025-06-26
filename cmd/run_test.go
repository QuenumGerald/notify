package cmd

import (
	"context"
	"testing"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func TestRun_Minimal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := &plugin.ExecutedCommand{}
	// Run should not panic or error with empty config
	err := Run(ctx, c)
	if err != nil {
		t.Errorf("Run failed: %v", err)
	}
}

func TestAutoRun_Minimal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := AutoRun(ctx)
	if err != nil {
		t.Errorf("AutoRun failed: %v", err)
	}
}
