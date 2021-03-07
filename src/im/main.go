package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"

	uuid "github.com/satori/go.uuid"
	"github.com/shuwenhe/shuwen-redis/src/redisConn"
)

var conn = redisConn.ConnectRedis()

func main() {
	user := CreateUser("123", "shuwen")
	fmt.Println("user = ", user)
}

func CreateUser(login, name string) string {
	llogin := strings.ToLower(login)
	lock := AcquireLockWithTimeout("user:"+llogin, 10, 10) // 加锁函数对小写的用户名进行加锁
	defer ReleaseLock("user:"+llogin, lock)
	if lock == "" {
		return ""
	}
	if conn.HGet("users:", llogin).Val() != "" {
		return ""
	}
	id := conn.Incr("user:id:").Val()
	pipeline := conn.TxPipeline()
	pipeline.HSet("users:", llogin, id)
	pipeline.HMSet(fmt.Sprintf("user:%s", strconv.Itoa(int(id))), "login", login, "id", id, "name", name, "followers", 0, "following", 0, "posts", 0, "sigup", time.Now().UnixNano())
	if _, err := pipeline.Exec(); err != nil {
		log.Println("pipeline err in CreateUser:", err)
		return ""
	}
	return strconv.Itoa(int(id))
}

func AcquireLockWithTimeout(lockname string, acquireTimeout, lockTimeout float64) string {
	identifier := uuid.NewV4().String()
	lockname = "lock" + lockname
	finallLockTimeout := math.Ceil(lockTimeout)

	end := time.Now().UnixNano() + int64(acquireTimeout*1e9)
	for time.Now().UnixNano() < end {
		if conn.SetNX(lockname, identifier, 0).Val() {
			conn.Expire(lockname, time.Duration(finallLockTimeout)*time.Second)
			return identifier
		} else if conn.TTL(lockname).Val() < 0 {
			conn.Expire(lockname, time.Duration(finallLockTimeout)*time.Second)
		}
		time.Sleep(10 * time.Millisecond)
	}
	return ""
}

func ReleaseLock(lockname, identifier string) bool {
	lockname = "lock:" + lockname
	var flag = true
	for flag {
		err := conn.Watch(func(tx *redis.Tx) error {
			pipe := tx.TxPipeline()
			fmt.Println(pipe)
			if tx.Get(lockname).Val() == identifier {
				pipe.Del(lockname)
				if _, err := pipe.Exec(); err != nil {
					return err
				}
				flag = true
				return nil
			}
			tx.Unwatch()
			flag = false
			return nil
		})
		if err != nil {
			log.Println("watch failed in ReleaseLock,err is:", err)
			return false
		}
	}
	return true
}
