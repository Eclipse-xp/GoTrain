package util

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var redisPool *redis.Pool
var mu sync.Mutex

func init() {
	if redisPool == nil {
		mu.Lock()
		if redisPool == nil {
			redisPool = newPool("localhost:6379")
		}
		mu.Unlock()
	}
}

func RedisPut(key string, value string) {
	conn := redisPool.Get()
	conn.Do("set", key, value)
	v2, _ := redis.String(conn.Do("get", key))
	fmt.Printf(v2)
	defer conn.Close()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		MaxActive:   10,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
