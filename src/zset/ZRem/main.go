package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZRem()
}

func ZRem() {
	res := conn.ZRem("zset-key", "c").Val() // 移除给定的成员,返回被移除成员的数量
	fmt.Println("res = ", res)
}
