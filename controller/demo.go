package controller

import (
  "web/ctx"

  "github.com/gin-gonic/gin"
)

func Get(ctx *ctx.New) {
  ctx.Success(gin.H{
    "data": "OK",
  })
}
