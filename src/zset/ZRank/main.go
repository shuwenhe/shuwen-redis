package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZRank()
}

func ZRank() {
	res := conn.ZRank("zset-key", "a").Val() // 排名
	fmt.Println("res-a = ", res)

	res2 := conn.ZRank("zset-key", "b").Val() // 排名
	fmt.Println("res-b = ", res2)

	res3 := conn.ZRank("zset-key", "c").Val() // 排名
	fmt.Println("res-c = ", res3)
}
