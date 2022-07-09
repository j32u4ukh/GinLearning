package database

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisDefaultPool *redis.Pool

// TODO: 這個腳本的存在意義在於，外部不需要直接操作 redis，可透過此腳本來封裝 redis 的各項功能
func init() {
	RedisDefaultPool = newPool("localhost:6379")
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func GetRedisConn() redis.Conn {
	return RedisDefaultPool.Get()
}

// TODO: 取值，若沒有，應自動存入或返回預設值
func GetRedis(conn redis.Conn, key string) ([]byte, error) {
	data, err := redis.Bytes(conn.Do("GET", key))
	return data, err
}
