package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZUnionStore()
}

func ZUnionStore() {
	res := conn.ZUnionStore("zset-u", &redis.ZStore{Aggregate: "min", Keys: []string{"zset-1", "zset-2"}}).Val()
	fmt.Println("res = ", res)
}
