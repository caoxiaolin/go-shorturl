package tests

import "testing"
import "github.com/caoxiaolin/go-shorturl/config"

func TestLoad(t *testing.T) {
	ret := config.Load()
	if ret.Database.Host != "127.0.0.1" {
		t.Error("Expected 127.0.0.1, got ", ret)
	}
}
