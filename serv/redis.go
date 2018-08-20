//service
package serv

import (
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/gomodule/redigo/redis"
)

var Rds redis.Conn

func init() {
	c, err := Conn()
	if err != nil {
		panic(err)
	}
	Rds = c
}

func Conn() (redis.Conn, error) {
	host := fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port)
	c, err := redis.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	if config.Cfg.Redis.Password != "" {
		if _, err := c.Do("AUTH", config.Cfg.Redis.Password); err != nil {
			c.Close()
			return nil, err
		}
	}
	if config.Cfg.Redis.Database > 0 {
		if _, err := c.Do("SELECT", config.Cfg.Redis.Database); err != nil {
			c.Close()
			return nil, err
		}
	}

	return c, nil
}
