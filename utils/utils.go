// 62进制转换规则
// '0'-'9' --- 0-9
// A-Z     --- 10-35
// a-z     --- 36-62

//utils tool
package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"math"
	"strconv"
)

//单个字符转数字，传入ascii
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

//数字转字符
func Int2str(i int) string {
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

//10进制转62进制
func Convert_10_to_62(num int) string {
	var res string
	var ys int
	s := num
	for s >= 1 {
		tmp := s / 62
		ys = s - tmp*62
		s = tmp
		res = Int2str(ys) + res
	}
	return res
}

//62进制转10进制
func Convert_62_to_10(str string) int64 {
	var res int64 = 0
	len := len(str)
	for k, v := range str {
		res = res + Str2int(v)*int64(math.Pow(float64(62), float64(len-1-k)))
	}
	return res
}

//获取短链接
func GetShortUrl(db *sql.DB, url string) string {
	var id int
	err := db.QueryRow(`INSERT INTO url (url, hits, create_time) VALUES ($1, 0, CURRENT_TIMESTAMP) RETURNING id`, url).Scan(&id)
	if err != nil {
		panic(err)
	}
	return Convert_10_to_62(int(id))
}

//获取原始链接
func GetOriUrl(db *sql.DB, url string) string {
	var oriurl string
	id := Convert_62_to_10(url)
	err := db.QueryRow(`SELECT url FROM url WHERE id = $1`, id).Scan(&oriurl)
	switch err {
	case sql.ErrNoRows:
		return ""
	case nil:
		//stmt, _ := db.Prepare("UPDATE url SET hits = hits + 1, last_access_time = NOW() WHERE id = ?")
	default:
		panic(err)
	}
	return oriurl
}
