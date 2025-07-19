package runner

import (
	"context"
	"notify/internal/config"
	"testing"
)

func TestRunner_Start_NoSubs(t *testing.T) {
	r := Runner{Subs: nil}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Should return immediately since no subs
	err := r.Start(ctx)
	if err != nil {
		t.Errorf("Runner.Start failed: %v", err)
	}
}

func TestRunner_Start_OneStdout(t *testing.T) {
	sub := config.Subscription{
		Name: "test",
		Node: "ws://localhost:26657",
		Query: "tm.event='NewBlock'",
		Sink: "stdout",
	}
	r := Runner{Subs: []config.Subscription{sub}}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		// Stop after a short delay
		<-ctx.Done()
	}()
	// Should not panic
	err := r.Start(ctx)
	if err != nil && err != context.Canceled {
		t.Errorf("Runner.Start failed: %v", err)
	}
}
