package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

type Client struct {
	Conn *redis.Client
}

func main() {
	var c *Client
	for i := 0; i < 3; i++ {
		go c.NotRans()
	}
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 3; i++ {
		go c.Trans()
	}
	time.Sleep(500 * time.Millisecond)
	defer c.Reset()
}

func (r *Client) Reset() {
	r.Conn.FlushDB()
}

func (r *Client) NotRans() {
	fmt.Println(r.Conn.Incr("notrans:").Val())
	time.Sleep(100 * time.Millisecond)
	fmt.Println(r.Conn.Decr("notrans:").Val())
}

func (r *Client) Trans() {
	pipeline := r.Conn.Pipeline()
	pipeline.Incr("trans:")
	time.Sleep(100 * time.Millisecond)
	pipeline.Decr("trans:")
	_, err := pipeline.Exec()
	if err != nil {
		log.Println("pipeline failed,the err is:", err)
	}
}
