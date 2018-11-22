package db

import (
  "time"
  "web/conf"

  "github.com/gomodule/redigo/redis"
)

var (
  RedisPool *redis.Pool
)

// 初始化redis连接池
func InitRedisPool() {
  if RedisPool == nil {
    RedisPool = &redis.Pool{
      MaxIdle:     3,
      IdleTimeout: 240 * time.Second,
      Dial: func() (redis.Conn, error) {
        conn, err := redis.Dial("tcp", conf.GetRedisdbUrl())
        return conn, err
      },
    }
  }
}

// 关闭redis连接池
func CloseRedisPool() {
  if RedisPool != nil {
    RedisPool.Close()
  }
}

// 从连接池中获取一个句柄
func GetRedis() redis.Conn {
  return RedisPool.Get()
}
