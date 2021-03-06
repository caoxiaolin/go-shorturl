//数据库相关的操作

//service
package serv

import (
	"database/sql"
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/caoxiaolin/go-shorturl/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.UserName, config.Cfg.Database.DbName, config.Cfg.Database.PassWord)
	db, _ = sql.Open("postgres", dsn)
	db.SetMaxOpenConns(config.Cfg.Database.MaxConn)
	db.Ping()
}

// GetShortUrl 将传入的URL入库并生成短链接返回
func GetShortUrl(url string) (string, error) {
	var id int64
	err := db.QueryRow(`INSERT INTO url (url, hits, create_time) VALUES ($1, 0, CURRENT_TIMESTAMP) RETURNING id`, url).Scan(&id)
	if err != nil {
		return "", err
	}
	return utils.Convert_10_to_62(int64(id)), nil
}

// GetOriUrl 根据传入的短链接查询获取原始链接并返回
func GetOriUrl(url string) (string, error) {
	var oriurl string
	id := utils.Convert_62_to_10(url)
	err := db.QueryRow(`SELECT url FROM url WHERE id = $1`, id).Scan(&oriurl)
	switch err {
	case sql.ErrNoRows:
		return "", nil
	case nil:
		stmt, err := db.Prepare("UPDATE url SET hits = hits + 1, last_access_time = NOW() WHERE id = $1")
		if err != nil {
			return "", err
		}
		_, err = stmt.Exec(id)
		if err != nil {
			return "", err
		}
	default:
		return "", err
	}
	return oriurl, nil
}
