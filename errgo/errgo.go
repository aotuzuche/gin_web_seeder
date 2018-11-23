package errgo

import (
  "errors"
  "net/http"
)

// 错误栈
type Stack []error

// 创建
func Create() *Stack {
  return new(Stack)
}

// 根据错误码换取错误信息
func GetErrorType(no interface{}) *errType {
  errStrNo := ""

  switch no.(type) {
  case int:
    errStrNo = no.(string)
  case error:
    errStrNo = no.(error).Error()
  }

  if errStrNo != "" && Error[errStrNo].Message != "" {
    err := Error[errStrNo]
    err.Code = errStrNo
    if err.Status == 0 {
      err.Status = http.StatusOK
    }
    return &err
  }

  err := Error[ErrServerError]
  err.Code = ErrServerError
  if err.Status == 0 {
    err.Status = http.StatusInternalServerError
  }

  return &err
}

// 添加一个错误进栈
func (s *Stack) PushError(errNo string) error {
  err := errors.New(errNo)
  *s = append(*s, err)
  return err
}

// 清空错误栈
func (s *Stack) ClearErrorStack() {
  *s = nil
}

// 弹出栈中的第一个错误(默认情况下弹出后就清空栈了)
func (s *Stack) PopError(clear ... bool) error {
  if len(*s) > 0 {
    first := (*s)[0]
    if clear != nil && clear[0] == false {
      *s = (*s)[1:]
    } else {
      s.ClearErrorStack()
    }
    return first
  }
  return nil
}
