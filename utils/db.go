// 62进制转换规则
// '0'-'9' --- 0-9
// A-Z     --- 10-35
// a-z     --- 36-62

//utils tool
package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// GetShortUrl 将传入的URL入库并生成短链接返回
func GetShortUrl(db *sql.DB, url string) string {
	var id int
    err := db.QueryRow(`INSERT INTO url (url, hits, create_time) VALUES ($1, 0, CURRENT_TIMESTAMP) RETURNING id`, url).Scan(&id)
    if err != nil {
        panic(err)
    }
	return Convert_10_to_62(int(id))
}

// GetOriUrl 根据传入的短链接查询获取原始链接并返回
func GetOriUrl(db *sql.DB, url string) string {
	var oriurl string
	id := Convert_62_to_10(url)
	err := db.QueryRow(`SELECT url FROM url WHERE id = $1`, id).Scan(&oriurl)
	switch err {
	case sql.ErrNoRows:
		return ""
	case nil:
		stmt, err := db.Prepare("UPDATE url SET hits = hits + 1, last_access_time = NOW() WHERE id = $1")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(id)
		if err != nil {
			panic(err)
		}
	default:
		panic(err)
	}
	return oriurl
}
