package main

import (
	"fmt"
	"sync"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var (
	conn = redisConn.ConnectRedis()
	wg   sync.WaitGroup
)

func main() {
	wg.Add(10)
	go LPop()
	wg.Wait()
}

func LPop() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		str := conn.LPop("list").Val()
		fmt.Printf("str %d = %s\n", i, str)
		str2 := conn.LPop("list2").Val()
		fmt.Printf("str2 %d = %s\n", i, str2)
	}
}
