package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	ZAdd()
}

func ZAdd() {
	res := conn.ZAdd("zset-1", &redis.Z{Member: "a", Score: 1}, &redis.Z{Member: "b", Score: 2}, &redis.Z{Member: "c", Score: 3}).Val()
	fmt.Println("res = ", res)

	res2 := conn.ZAdd("zset-2", &redis.Z{Member: "b", Score: 4}, &redis.Z{Member: "d", Score: 0}, &redis.Z{Member: "c", Score: 1}).Val()
	fmt.Println("res = ", res2)
}
