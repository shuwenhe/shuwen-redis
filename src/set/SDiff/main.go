package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SDiff()
}

func SDiff() {
	set := conn.SDiff("skey1", "skey2").Val() // key1=[a,b,c] key2=[b] => [a,c]
	for k, v := range set {
		fmt.Printf("set%d = %s\n", k, v)
	}
}
