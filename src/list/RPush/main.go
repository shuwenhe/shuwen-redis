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
	conn.RPush("list", "item1")
	res := conn.RPush("list", "item2").Val()
	fmt.Println("res = ", res)

	conn.RPush("list2", "item3")
	res2 := conn.RPush("list2", "item4").Val()
	fmt.Println("res2 = ", res2)
}
