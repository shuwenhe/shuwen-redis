package redisConn

import (
	"log"
	"shuwen-redis/src/config"

	"github.com/go-redis/redis/v7"
)

func ConnectRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := conn.Ping().Result(); err != nil {
		log.Fatalf("Connect to redis client failed,err:%v\n", err)
	}
	return conn
}
