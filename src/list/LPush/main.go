package main

import (
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	// LPush()
}

func LPush() {
	conn.LPush("list-key", "first")
}
