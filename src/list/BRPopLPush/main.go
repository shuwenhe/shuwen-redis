package main

import (
	"fmt"
	"time"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	BRPopLPush()
}

func BRPopLPush() { // 将一个元素从一个列表移动到另一个列表，并返回被移动的元素
	str := conn.BRPopLPush("list2", "list", 1*time.Second).Val()
	fmt.Println("str = ", str)
}
