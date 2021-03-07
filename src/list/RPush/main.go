package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	RPush()
}

func RPush() {
	res := conn.RPush("sort-input", 1, 9, 3, 7, 5).Val()
	fmt.Println("res = ", res)
}
