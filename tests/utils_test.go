package tests

import "testing"
import "github.com/caoxiaolin/go-shorturl/utils"

func TestConvert_10_to_62(t *testing.T) {
	ret := utils.Convert_10_to_62(10000)
	if ret != "2bI" {
		t.Error("Expected 2bI, got ", ret)
	}
}

func TestConvert_62_to_10(t *testing.T) {
	ret := utils.Convert_62_to_10("2bI")
	if ret != 10000 {
		t.Error("Expected 10000, got ", ret)
	}
}
