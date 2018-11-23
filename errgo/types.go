package errgo

import "net/http"

// 错误数据的结构
type errType struct {
  Message string
  Status  int
  Code    string
}

const (
  // 系统级错误 200xxx
  ErrIdError    = "200000"
  ErrSkipRange  = "200001"
  ErrLimitRange = "200002"
  ErrForbidden  = "200003"
  ErrNeedLogin  = "200008"

  // 默认错误
  ErrServerError = "999999"
)

// 错误列表
var Error = map[string]errType{
  ErrIdError:    {Message: "非法的id"},
  ErrSkipRange:  {Message: "skip取值范围错误"},
  ErrLimitRange: {Message: "limit取值范围错误"},
  ErrForbidden:  {Message: "权限不足", Status: http.StatusForbidden},
  ErrNeedLogin:  {Message: "请重新登录", Status: http.StatusForbidden},

  ErrServerError: {Message: "系统错误", Status: http.StatusInternalServerError},
}
