package main

import (
  "flag"
  "web/conf"
  "web/db"
  "web/router"

  "github.com/gin-gonic/gin"
  "github.com/zouyx/agollo"
)

func init() {
  // 连接apollo
  agollo.InitCustomConfig(conf.GetApolloConfig)
  agollo.Start()

  // 连接数据库
  db.ConnectMgoDB()
  db.InitRedisPool()
}

func main() {
  // 退出时关闭数据库连接
  defer db.CloseMgoDB()
  defer db.CloseRedisPool()

  // 初始化gin框架
  gin.SetMode(gin.DebugMode)

  // Default With the Logger and Recovery middleware already attached
  g := gin.Default() // gin.New()

  // 注册路由
  router.Register(g)

  // 获取端口号
  // e.g.  go run main.go --port=9090
  port := ""
  flag.StringVar(&port, "port", "9090", "port addr")
  flag.Parse()

  // 启动
  g.Run(":" + port)
}
