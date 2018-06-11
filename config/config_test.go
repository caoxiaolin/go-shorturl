package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := Load("./")
	if cfg.Database.Host != "127.0.0.1" {
		t.Error("Expected 127.0.0.1, but got ", cfg.Database.Host)
	}
}
