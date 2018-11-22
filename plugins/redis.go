package plugins

import (
  "web/db"
  "web/util"

  "github.com/gomodule/redigo/redis"
)

func GetRedisConn() redis.Conn {
  rds := db.GetRedis()
  if rds != nil {
    util.Println("[RDS] 😄 OK")
  }
  return rds
}

func CloseRedisConn(r redis.Conn) {
  if r != nil {
    r.Close()
    util.Println("[RDS] 👋 CLOSED")
  }
}