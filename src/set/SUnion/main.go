package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SUnion()
}

func SUnion() {
	set := conn.SUnion("skey1", "skey2").Val()
	for k, v := range set {
		fmt.Printf("v%d = %s\n", k, v)
	}
}
