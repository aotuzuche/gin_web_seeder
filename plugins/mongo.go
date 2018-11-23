package plugins

import (
  "web/db"
  "web/util"

  "gopkg.in/mgo.v2"
)

func GetMgoSession() (*mgo.Database, func()) {
  // 创建mgodb的session
  mg, closer, err := db.GetMgoSession()
  if err != nil {
    util.Println("[MGO] 😈 Error")
    panic(err)
  }
  if mg != nil {
    util.Println("[MGO] 😄 OK")
  }
  return mg, closer
}

func CloseMgoSession(f func()) {
  if f != nil {
    f()
    util.Println("[MGO] 👋 CLOSED")
  }
}