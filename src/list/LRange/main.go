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
	res := conn.LRange("list", 0, -1).Val()
	for k, v := range res {
		fmt.Printf("res %d = %s\n", k, v)
	}

	res2 := conn.LRange("list2", 0, -1).Val()
	for k, v := range res2 {
		fmt.Printf("res2 %d = %s\n", k, v)
	}
}
