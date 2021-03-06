package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SAdd()
}

func SAdd() {
	res := conn.SAdd("skey1", "a", "b", "c").Val()
	fmt.Println("res= ", res)
}
