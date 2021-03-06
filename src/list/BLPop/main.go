package main

import (
	"fmt"
	"time"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	BLPop()
}

func BLPop() { //BLPop从左到右检查传入的列表，并对最先遇到的非空列表执行弹出操作
	str := conn.BLPop(1*time.Second, "list", "list2").Val()
	for k, v := range str {
		fmt.Printf("v%d = %s\n", k, v)
	}
}
