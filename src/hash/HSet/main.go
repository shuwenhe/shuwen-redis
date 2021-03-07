package main

import (
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	HSet()
}

func HSet() {
	conn.HSet("d-7", "field", 5)
	conn.HSet("d-15", "field", 1)
	conn.HSet("d-23", "field", 9)
	conn.HSet("d-110", "field", 3)
}
