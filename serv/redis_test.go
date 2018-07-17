package serv

import (
	"testing"
    "github.com/gomodule/redigo/redis"
)

func TestRedis(t *testing.T){
    rdskey := "test_redis"
    Rds.Do("SET", rdskey, "this is a test")
    if rdsval, _ := redis.String(Rds.Do("GET", rdskey)); rdsval != "this is a test" {
        t.Error("Expected this is a test, but got", rdsval)
    }
}
