package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZScore()
}

func ZScore() {
	res := conn.ZScore("zset-key", "c").Val()
	fmt.Println("res = ", res)
}
