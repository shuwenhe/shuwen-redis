package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZInterStore()
}

func ZInterStore() {
	res := conn.ZInterStore("zset-i", &redis.ZStore{Keys: []string{"ZSet-1", "ZSet-2"}}).Val()
	fmt.Println("res = ", res)
}
