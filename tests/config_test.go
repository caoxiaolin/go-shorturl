package tests

import (
	"github.com/caoxiaolin/go-shorturl/config"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := config.Load()
	if cfg.Database.Host != "127.0.0.1" {
		t.Error("Expected 127.0.0.1, got ", cfg.Database.Host)
	}
}
