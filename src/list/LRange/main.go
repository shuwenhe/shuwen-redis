package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	LRange()
}

func LRange() {
	res := conn.LRange("list-key", 0, -1).Val()
	for k, v := range res {
		fmt.Printf("v%d = %s\n", k, v)
	}
}
