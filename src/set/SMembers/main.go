package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SMembers()
}

func SMembers() {
	res := conn.SMembers("skey1").Val()
	for k, v := range res {
		fmt.Printf("setVal%d = %s\n", k, v)
	}
}
