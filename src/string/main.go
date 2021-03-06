package main

import (
	"fmt"
	"shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	// Get()
	// Set()
	// Incr()
	// IncrBy()
	// Decr()
	// DecrBy()

}

func connRedis() {
	fmt.Println("conn = ", conn) // Redis<62.234.11.179:6379 db:15>
}

func Set() {
	res := conn.Set("num", 9, 0).Val()
	fmt.Println("res = ", res)
}

func Get() {
	res := conn.Get("num")
	fmt.Println("num=", res)
}

func Incr() { // incr自增
	res := conn.Incr("num").Val()
	fmt.Println("res = ", res)
}

func IncrBy() {
	res := conn.IncrBy("num", 3)
	fmt.Println("res = ", res)
}

func Decr() {
	res := conn.Decr("num")
	fmt.Println("res = ", res)
}

func DecrBy() {
	res := conn.DecrBy("num", 5)
	fmt.Println("res = ", res)
}
