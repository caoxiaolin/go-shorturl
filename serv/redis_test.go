package serv

import (
	"github.com/caoxiaolin/go-shorturl/config"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	rdskey := "test_redis"
	Rds.Do("SET", rdskey, "this is a test")
	if rdsval, _ := redis.String(Rds.Do("GET", rdskey)); rdsval != "this is a test" {
		t.Error("Expected this is a test, but got", rdsval)
	}
}

func TestConn(t *testing.T) {
	config.Cfg.Redis.Port = 1000
	_, err := Conn()
	if err == nil {
		t.Error(err)
	}

	config.Cfg.Redis.Port = 6379
	config.Cfg.Redis.Password = "123456"
	_, err = Conn()
	if err == nil {
		t.Error(err)
	}

	config.Cfg.Redis.Password = ""
	config.Cfg.Redis.Database = 100
	_, err = Conn()
	if err == nil {
		t.Error(err)
	}
}
