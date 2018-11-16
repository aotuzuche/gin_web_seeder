package util

import (
  "fmt"

  "github.com/gin-gonic/gin"
)

// 打印的封装，只在debug状态打印
func Println(a ... interface{}) {
  if gin.IsDebugging() {
    fmt.Println(a...)
  }
}
