package main

import (
	"fmt"
	"time"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	Expire()
}

func Expire() {
	conn.Set("key", "value", 0)
	res := conn.Get("key").Val()
	fmt.Println("res = ", res)
	conn.Expire("key", 1*time.Second)
	time.Sleep(2 * time.Second)
	conn.Get("key").Val()
	conn.Set("key", "value2", 0)
	conn.Expire("key", 100*time.Second)
}
