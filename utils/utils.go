// 62进制转换规则
// '0'-'9' --- 0-9
// A-Z     --- 10-35
// a-z     --- 36-62

//utils tool
package utils

import (
	"math"
	"net/http"
	"strconv"
)

// Str2int用来单个字符转数字，传入ascii
func Str2int(i rune) int64 {
	var res int32
	if i >= 48 && i <= 57 {
		res = i - 48
	} else if i >= 65 && i <= 90 {
		res = i - 65 + 10
	} else {
		res = i - 97 + 36
	}
	return int64(res)
}

// Int2str用来数字转字符
func Int2str(i int64) string {
	var res string
	if i >= 0 && i <= 9 {
		res = strconv.Itoa(int(i))
	} else if i >= 10 && i <= 35 {
		res = string(i - 10 + 'A')
	} else {
		res = string(i - 36 + 'a')
	}
	return res
}

// Convert_10_to_62用于10进制转62进制
func Convert_10_to_62(num int64) string {
	var res string
	var ys int64
	s := num
	for s >= 1 {
		tmp := s / 62
		ys = s - tmp*62
		s = tmp
		res = Int2str(ys) + res
	}
	return res
}

// Convert_62_to_10用于62进制转10进制
func Convert_62_to_10(str string) int64 {
	var res int64
	len := len(str)
	for k, v := range str {
		res = res + Str2int(v)*int64(math.Pow(float64(62), float64(len-1-k)))
	}
	return res
}

// GetPostUrl 从http请求中解析并校验url
func GetPostUrl(r *http.Request) string {
	r.ParseForm()
	if r.Form["url"] == nil {
		return ""
	}
	return r.Form["url"][0]
}
