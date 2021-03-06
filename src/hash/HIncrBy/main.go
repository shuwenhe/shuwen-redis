package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HIncry()
}

func HIncry() {
	res := conn.HIncrBy("hash-key2", "num", 1).Val()
	fmt.Println("res = ", res)
}
