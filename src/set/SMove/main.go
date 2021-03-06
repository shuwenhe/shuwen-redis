package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SMove()
}

func SMove() {
	b := conn.SMove("set-key", "set-key2", "b").Val() // 从key->key2
	fmt.Println("b = ", b)
}
