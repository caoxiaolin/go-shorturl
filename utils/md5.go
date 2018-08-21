package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	if str == "" {
		return ""
	}
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}
