package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZCard()
}

func ZCard() {
	res := conn.ZCard("zset-key").Val() // 有序集合成员数量
	fmt.Println("res = ", res)
}
