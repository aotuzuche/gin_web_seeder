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
  if !conf.UseRedis {
    return
  }

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

func CloseRedisPool() {
  if RedisPool != nil {
    RedisPool.Close()
  }
}

func GetRedis() redis.Conn {
  if !conf.UseRedis {
    return nil
  }
  return RedisPool.Get()
}
