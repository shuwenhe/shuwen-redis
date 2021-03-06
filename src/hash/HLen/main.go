package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HLen()
}

func HLen() {
	res := conn.HLen("hash-key").Val()
	fmt.Println("The len of the hash = ", res)
}
