package main

import (
	"fmt"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	// RPush()
	// RPush2()
	// RPush3()
}

func RPush() {
	conn.RPush("list-key", "last")
}

func RPush2() {
	res := conn.RPush("list-key", "new last").Val()
	fmt.Println("res = ", res)
}

func RPush3() {
	res := conn.RPush("list-key", "a", "b", "c").Val() // values => ...interface{}
	fmt.Println("res = ", res)
}
