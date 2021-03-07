package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	go RunPubSub()
	Publish(6)
	time.Sleep(1 * time.Second)
	defer Reset()
}

func RunPubSub() { // 启动发送者线程
	pubSub := conn.Subscribe("channel")
	defer pubSub.Close()
	var count int32 = 0
	for item := range pubSub.Channel() { // 通过遍历Channel的执行结果来监听订阅消息
		fmt.Println(item.String()) // 打印接收到的每条消息
		atomic.AddInt32(&count, 1) // 接收一条订阅反馈消息
		fmt.Println(count)         // 3条发布者发送的消息
		switch count {
		case 4: // 执行退出操作，停止监听新消息
			if err := pubSub.Unsubscribe("channel"); err != nil {
				log.Println("unsubscribe faile, err:", err)
			} else {
				fmt.Println("unsubscribe success")
			}
		case 5: // 客户端在接收到退订反馈消息之后就不再接收消息
			break
		default:
		}
	}
}

func Publish(n int) {
	time.Sleep(1 * time.Second) // 休眠，让订阅者有足够的时间连接服务器并监听消息
	for n > 0 {
		conn.Publish("channel", n) // 消息一条一条出现
		n--
	}
}

func Reset() {
	conn.FlushDB()
}
