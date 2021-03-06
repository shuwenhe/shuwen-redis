package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HMGet()
}

func HMGet() {
	res := conn.HMGet("hash-key", "k2", "k3").Val()
	for _, v := range res {
		fmt.Println("v = ", v)
	}
}
