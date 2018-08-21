package utils

import (
	"testing"
)

/**
 * 测试十进制转62进制
 */
func TestMd5(t *testing.T) {
	ret := MD5("1")
	if ret != "c4ca4238a0b923820dcc509a6f75849b" {
		t.Error("Expected c4ca4238a0b923820dcc509a6f75849b, but got", ret)
	}
}
