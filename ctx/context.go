package ctx

import (
  "net/http"
  "web/errgo"
  "web/plugins"
  "web/util"

  "github.com/gin-gonic/gin"
)

type New struct {
  Ctx     *gin.Context
  RawData []byte
  Errgo   *errgo.Stack
  plugins.Plugins
}

func CreateCtx(fn func(*New)) func(*gin.Context) {
  return func(c *gin.Context) {
    util.Println()
    util.Println("------------------------------------------")
    util.Println()

    // 创建上下文
    bytes, _ := c.GetRawData()
    ctx := &New{
      c,
      bytes,
      errgo.Create(),
      plugins.CreatePlugins(),
    }

    // defer
    defer plugins.DestroyPlugins(ctx.Plugins)

    // 调用控制器
    fn(ctx)
  }
}

// 成功处理
func (c *New) Success(data gin.H) {
  resp := gin.H{
    "msg":  "ok",
    "code": 0,
  }

  if len(data) > 1 { // Almost the length is more than 1, so just check it first.
    resp["data"] = data
  } else if data["data"] != nil {
    resp["data"] = data["data"]
  } else if data != nil && len(data) > 0 {
    resp["data"] = data
  }

  status := http.StatusOK
  if data == nil {
    status = http.StatusNoContent
  }

  c.Ctx.JSON(status, resp)
}

// 处理错误
func (c *New) Error(errNo interface{}) {

  // 根据错误号获取错误内容（错误号是个int或error）
  err := errgo.GetErrorType(errNo)

  util.Println()
  util.Println(" >>> ORIGIN:", errNo)
  util.Println(" >>> ERROR:", err.Message)
  util.Println(" >>> ERROR CODE:", err.Code)
  util.Println(" >>> REQUEST METHOD:", c.Ctx.Request.Method)
  util.Println(" >>> REQUEST URL:", c.Ctx.Request.URL.String())
  util.Println(" >>> USER AGENT:", c.Ctx.Request.UserAgent())
  util.Println()

  c.Ctx.JSON(err.Status, gin.H{
    "msg":  err.Message,
    "code": err.Code,
    "data": nil,
  })
}

// 响应，如果有错误走Error，否则走Success，但为了让success里的数据懒执行，需要用一个func包装
func (c *New) Response(err interface{}, succ func() gin.H) {
  if err == nil {
    c.Success(succ())
    return
  }
  c.Error(err)
}
