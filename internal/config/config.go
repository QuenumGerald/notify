package config

import (
	"os"
	"os/user"
	"path/filepath"
	"gopkg.in/yaml.v3"
)

type Subscription struct {
	Name    string `yaml:"name"`
	Node    string `yaml:"node"`
	Query   string `yaml:"query"`
	Sink    string `yaml:"sink"`
	Webhook string `yaml:"webhook"`
}

func GetConfigPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(u.HomeDir, ".ignite")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "notify.yaml"), nil
}

func LoadSubscriptions(file string) ([]Subscription, error) {
	f, err := os.Open(file)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var subs []Subscription
	if err := yaml.NewDecoder(f).Decode(&subs); err != nil && err.Error() != "EOF" {
		return nil, err
	}
	return subs, nil
}

func SaveSubscriptions(file string, subs []Subscription) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewEncoder(f).Encode(subs)
}
