package main

import (
	"context"
	"testing"
	"github.com/ignite/cli/v29/ignite/services/plugin"
)

func TestManifest(t *testing.T) {
	app := notifyApp{}
	_, err := app.Manifest(context.Background())
	if err != nil {
		t.Errorf("Manifest failed: %v", err)
	}
}

func TestExecute_Unknown(t *testing.T) {
	app := notifyApp{}
	err := app.Execute(context.Background(), &plugin.ExecutedCommand{Path: "notify unknown"}, nil)
	if err == nil {
		t.Error("Expected error for unknown command path")
	}
}

func TestExecuteHookPre(t *testing.T) {
	app := notifyApp{}
	err := app.ExecuteHookPre(context.Background(), nil, nil)
	if err != nil {
		t.Errorf("ExecuteHookPre failed: %v", err)
	}
}

func TestExecuteHookPost(t *testing.T) {
	app := notifyApp{}
	h := &plugin.ExecutedHook{}
	err := app.ExecuteHookPost(context.Background(), h, nil)
	if err != nil {
		t.Errorf("ExecuteHookPost failed: %v", err)
	}
}
