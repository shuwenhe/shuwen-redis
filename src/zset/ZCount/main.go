package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZCount()
}

func ZCount() {
	res := conn.ZCount("zset-key", "0", "3").Val() // (min,max)
	fmt.Println("res = ", res)
}
