package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	RPop()
}

func RPop() {
	str := conn.RPop("list-key").Val()
	fmt.Println("str = ", str)
}
