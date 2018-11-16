package router

import (
  "web/controller"
  "web/ctx"

  "github.com/gin-gonic/gin"
)

func registerDemoRouter(g *gin.RouterGroup) {

  // 这是一个测试请求
  g.GET("", /*middleware.Jwt,*/ ctx.CreateCtx(controller.Get))

}
