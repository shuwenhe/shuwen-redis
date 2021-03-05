package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "62.234.11.179:6379",
		Password: "shuwen123456",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connet redis failed,err:%v\n", err)
		return
	}
	fmt.Println("redis connect success!")
}
