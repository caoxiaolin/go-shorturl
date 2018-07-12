//service
package serv

import (
	"fmt"
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/gomodule/redigo/redis"
)

var Rds redis.Conn

func init() {
	host := fmt.Sprintf("%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port)
	c, err := redis.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	if config.Cfg.Redis.Password != "" {
		if _, err := c.Do("AUTH", config.Cfg.Redis.Password); err != nil {
			c.Close()
			panic(err)
		}
	}
	if config.Cfg.Redis.Database > 0 {
		if _, err := c.Do("SELECT", config.Cfg.Redis.Database); err != nil {
			c.Close()
			panic(err)
		}
	}
	Rds = c
}
