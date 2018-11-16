package router

import "github.com/gin-gonic/gin"

func Register(g *gin.Engine) {

  // 路由分组
  registerDemoRouter(g.Group("/demo"))

}
