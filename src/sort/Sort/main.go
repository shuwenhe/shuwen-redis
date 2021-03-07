package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	SortList()
	SortHash()
}

func SortList() {
	res := conn.Sort("sort-input", &redis.Sort{Order: "ASC"}).Val()
	fmt.Println("res = ", res)
}

func SortHash() {
	res := conn.Sort("sort-input", &redis.Sort{Alpha: true}).Val()
	res = conn.Sort("sort-input", &redis.Sort{By: "d-*->field"}).Val()
	res = conn.Sort("sort-input", &redis.Sort{By: "d-*->field", Get: []string{"d-*->field"}}).Val()
	fmt.Println("res = ", res)
}
