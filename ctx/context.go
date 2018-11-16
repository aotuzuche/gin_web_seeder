package ctx

import (
  "net/http"
  "web/db"
  "web/errgo"
  "web/util"

  "github.com/gin-gonic/gin"
  "github.com/gomodule/redigo/redis"
  "gopkg.in/mgo.v2"
)

type New struct {
  Ctx         *gin.Context
  RawData     []byte
  MgoDB       *mgo.Database
  Redis       redis.Conn
  Errgo       *errgo.Stack
  mgoDBCloser func()
}

func CreateCtx(fn func(*New)) func(*gin.Context) {
  return func(c *gin.Context) {
    util.Println()
    util.Println("------------------------------------------")
    util.Println()

    // 创建上下文
    ctx, err := NewCtx(c)

    // 如果创建过程中有报错，返回错误
    if err != nil {
      ctx.Error(err)
      return
    }

    // defer
    defer ctx.Close()

    // 调用控制器
    fn(ctx)
  }
}

// 创建上下文，连接mgo与redis数据库
func NewCtx(c *gin.Context) (*New, error) {
  bytes, _ := c.GetRawData()

  mg, closer, err := db.CloneMgoDB()
  if err != nil {
    util.Println("[MGO] 😈 Error")
    return nil, err
  }
  if mg != nil {
    util.Println("[MGO] 😄 OK")
  }

  rds := db.GetRedis()
  if rds != nil {
    util.Println("[RDS] 😄 OK")
  }

  return &New{
    c,
    bytes,
    mg,
    rds,
    errgo.Create(),
    closer,
  }, nil
}

// 创建不连接数据库的上下文
func NewBaseCtx(c *gin.Context) *New {
  bytes, _ := c.GetRawData()
  return &New{
    Ctx:     c,
    RawData: bytes,
    Errgo:   errgo.Create(),
  }
}

// 关闭数据库连接
func (c *New) Close() {
  if c.mgoDBCloser != nil {
    c.mgoDBCloser()
    util.Println("[MGO] 👋 CLOSED")
  }
  if c.Redis != nil {
    c.Redis.Close()
    util.Println("[RDS] 👋 CLOSED")
  }
}

// 成功处理
func (c *New) Success(data gin.H) {
  respH := gin.H{
    "msg":  "ok",
    "code": 0,
  }

  if len(data) > 1 { // Almost the length is more than 1, so just check it first.
    respH["data"] = data
  } else if data["data"] != nil {
    respH["data"] = data["data"]
  } else if data != nil && len(data) > 0 {
    respH["data"] = data
  }

  status := http.StatusOK

  if data == nil {
    status = http.StatusNoContent
  }

  c.Ctx.JSON(status, respH)
}

// 处理错误
func (c *New) Error(errNo interface{}) {

  // 根据错误号获取错误内容（错误号是个int或error）
  err := errgo.Get(errNo)

  util.Println()
  util.Println(" >>> ERROR:", err.Message)
  util.Println(" >>> ERROR CODE:", err.Code)
  util.Println(" >>> REQUEST METHOD:", c.Ctx.Request.Method)
  util.Println(" >>> REQUEST URL:", c.Ctx.Request.URL.String())
  util.Println(" >>> USER AGENT:", c.Ctx.Request.UserAgent())
  util.Println()

  // 清除错误栈
  c.Errgo.ClearErrorStack()

  c.Ctx.JSON(err.Status, gin.H{
    "msg":  err.Message,
    "code": err.Code,
    "data": nil,
  })
}

// 响应，如果有错误走Error，否则走Success
func (c *New) Response(err interface{}, succ gin.H) {
  if err == nil {
    c.Success(succ)
    return
  }
  c.Error(err)
}
