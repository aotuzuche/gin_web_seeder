package plugins

import (
  "longRentServer/util"
  "web/db"
  "web/errgo"

  "github.com/gomodule/redigo/redis"
  "gopkg.in/mgo.v2"
)

// ctx的插件(插件名不允许叫Ctx、RawData或Plugins)
type Plugins struct {
  MgoDB     *mgo.Database
  Redis     redis.Conn
  Errgo     *errgo.Stack
  mgoCloser func()
}

// 在发生请求时，初始化插件
func CreatePlugins() Plugins {
  // 创建mgodb的session
  mg, closer, err := db.CloneMgoDB()
  if err != nil {
    util.Println("[MGO] 😈 Error")
  }
  if mg != nil {
    util.Println("[MGO] 😄 OK")
  }

  // 从redis的连接池中取一个
  rds := db.GetRedis()
  if rds != nil {
    util.Println("[RDS] 😄 OK")
  }

  // 返回插件集
  return Plugins{
    MgoDB: mg,
    Redis: rds,
    Errgo: errgo.Create(),
    mgoCloser: closer,
  }
}

// 在请求结束时，做收尾处理
func DestroyPlugins(p Plugins) {
  // 关闭mgodb的连接
  if p.mgoCloser != nil {
    p.mgoCloser()
    util.Println("[MGO] 👋 CLOSED")
  }
  // 关闭redis连接
  if p.Redis != nil {
    p.Redis.Close()
    util.Println("[RDS] 👋 CLOSED")
  }
}
