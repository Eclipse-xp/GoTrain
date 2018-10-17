package util

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func RedisPut(key string, value string) {
	conn, _ := redis.Dial("tcp", "localhost:6379")
	conn.Do("set", key, value)
	v2, _ := redis.String(conn.Do("get", "redisGoStr"))
	fmt.Printf(v2)
	defer conn.Close()
}
