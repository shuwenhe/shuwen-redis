package main

import (
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	LTrim()
}

func LTrim() {
	conn.LTrim("list-key", 2, -1)
}
