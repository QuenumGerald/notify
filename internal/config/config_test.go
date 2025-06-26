package config

import (
	"os"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	path, err := GetConfigPath()
	if err != nil {
		t.Fatalf("GetConfigPath failed: %v", err)
	}
	if path == "" {
		t.Fatal("Config path should not be empty")
	}
}

func TestSaveAndLoadSubscriptions(t *testing.T) {
	file := "test_notify.yaml"
	subs := []Subscription{{Name: "foo", Node: "ws://localhost:26657", Query: "tm.event='NewBlock'", Sink: "stdout", Webhook: ""}}
	if err := SaveSubscriptions(file, subs); err != nil {
		t.Fatalf("SaveSubscriptions failed: %v", err)
	}
	loaded, err := LoadSubscriptions(file)
	if err != nil {
		t.Fatalf("LoadSubscriptions failed: %v", err)
	}
	if len(loaded) != 1 || loaded[0].Name != "foo" {
		t.Fatalf("Loaded subscriptions do not match: %+v", loaded)
	}
	os.Remove(file)
}
