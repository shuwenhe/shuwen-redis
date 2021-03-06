package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SInter()
}

func SInter() {
	set := conn.SInter("skey1", "skey2").Val()
	for k, v := range set {
		fmt.Printf("v%d = %s\n", k, v)
	}
}
