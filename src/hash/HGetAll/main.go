package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HGetAll()
}

func HGetAll() {
	res := conn.HGetAll("hash-key").Val()
	fmt.Println("res = ", res)
}
