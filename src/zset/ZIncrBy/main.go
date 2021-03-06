package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZIncrBy()
}

func ZIncrBy() {
	res := conn.ZIncrBy("zset-key", 3, "c").Val()
	fmt.Println("res = ", res)
}
