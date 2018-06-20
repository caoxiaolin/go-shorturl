package utils

import (
	"testing"
)

/**
 * 测试十进制转62进制
 */
func TestConvert_10_to_62(t *testing.T) {
	ret := Convert_10_to_62(10000)
	if ret != "2bI" {
		t.Error("Expected 2bI, but got ", ret)
	}
}

/**
 * 测试62进制转10进制
 */
func TestConvert_62_to_10(t *testing.T) {
	ret := Convert_62_to_10("2bI")
	if ret != 10000 {
		t.Error("Expected 10000, but got ", ret)
	}
}
