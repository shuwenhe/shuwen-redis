package model

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type Client struct {
	Conn *redis.Client
}

// CheckToken 尝试获取并返回令牌对应的用户
func (r *Client) CheckToken(token string) string {
	return r.Conn.HGet("login:", token).Val()
}

// UpdateToken 更新令牌
func (r *Client) UpdateToken(token, user, item string) {
	timestamp := time.Now().Unix()
	r.Conn.HSet("login:", token, user)
	r.Conn.ZAdd("recent", &redis.Z{Score: float64(timestamp), Member: token})
	if item != "" {
		r.Conn.ZAdd("viewed:"+token, item, float64(timestamp))
		r.Conn.ZRemRangeByRank("viewed:"+token, 0, -26)
	}
}
