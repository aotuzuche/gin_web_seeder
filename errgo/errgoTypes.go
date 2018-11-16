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
  ErrIdError:    {"非法的id", http.StatusOK, ""},
  ErrSkipRange:  {"skip取值范围错误", http.StatusInternalServerError, ""},
  ErrLimitRange: {"limit取值范围错误", http.StatusInternalServerError, ""},
  ErrForbidden:  {"权限不足", http.StatusForbidden, ""},
  ErrNeedLogin:  {"请重新登录", http.StatusForbidden, ""},

  ErrServerError: {"系统错误", http.StatusInternalServerError, ""},
}
