package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	// LPop()
}

func LPop() {
	str := conn.LPop("list-key").Val()
	fmt.Println("str = ", str)
}
