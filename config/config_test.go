package config

import (
	"testing"
)

/**
 * 测试获取配置项
 */
func TestConfig(t *testing.T) {
    ret := Cfg.Server.Host
	if ret != "127.0.0.1" {
		t.Error("Expected 127.0.0.1, but got", ret)
	}
}
