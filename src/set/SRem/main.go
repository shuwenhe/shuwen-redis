package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SRem()
}

func SRem() {
	res := conn.SRem("skey2", "a", "b", "c", "d").Val()
	fmt.Println("The num of remove set = ", res)
}
