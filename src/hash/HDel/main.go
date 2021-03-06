package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HDel()
}

func HDel() {
	res := conn.HDel("hash-key", "k2", "k3").Val()
	fmt.Println("The len of the hash = ", res)
}
