package plugins

import (
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
  mg, closer := CreateMgoSession()

  // 返回插件集
  return Plugins{
    MgoDB:     mg,
    Redis:     GetRedisConn(),
    Errgo:     errgo.Create(),
    mgoCloser: closer,
  }
}

// 在请求结束时，做收尾处理
func DestroyPlugins(p Plugins) {
  // 关闭mgodb的连接
  CloseMgoSession(p.mgoCloser)
  // 关闭redis连接
  CloseRedisConn(p.Redis)
}
