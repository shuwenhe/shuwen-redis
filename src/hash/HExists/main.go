package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HExists()
}

func HExists() {
	b := conn.HExists("hash-key2", "short").Val()
	fmt.Println("b = ", b)
}
