package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZRangeWithScores()
}

func ZRangeWithScores() {
	res := conn.ZRangeWithScores("zset-u", 0, -1).Val() // 移除给定的成员,返回被移除成员的数量
	for _, v := range res {
		fmt.Println("v=", v)
	}
}
