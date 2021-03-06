package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HKeys()
}

func HKeys() {
	res := conn.HKeys("hash-key2").Val()
	fmt.Println("res = ", res)
}
