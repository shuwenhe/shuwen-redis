package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HMSet()
}

func HMSet() {
	b := conn.HMSet("hash-key2", map[string]interface{}{
		"short": "hello",
		"long":  "1000",
	}).Val()
	fmt.Println("b = ", b)
}
