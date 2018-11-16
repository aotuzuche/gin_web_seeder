package errgo

import (
  "errors"
  "time"

  "gopkg.in/mgo.v2/bson"
)

// 错误栈
type Stack []error

// 根据错误码换取错误信息
func Get(no interface{}) *errType {
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
    return &err
  }

  err := Error[ErrServerError]
  err.Code = ErrServerError

  return &err
}

// 创建
func Create() *Stack {
  return new(Stack)
}

// 判断bool返回值
func (s *Stack) True(bool bool, errNo string) error {
  if bool {
    return s.Error(errNo)
  }
  return nil
}

// 判断func返回值
func (s *Stack) FuncTrue(fn func() bool, errNo string) error {
  return s.True(fn(), errNo)
}

// 判断int是否小于一个值
func (s *Stack) IntLessThen(val int, min int, errNo string) error {
  return s.True(val < min, errNo)
}

// 判断int是否大于一个值
func (s *Stack) IntMoreThen(val int, max int, errNo string) error {
  return s.True(val > max, errNo)
}

// 判断一个值是否为objectId(用mongodb时会用到)
func (s *Stack) StringNotObjectId(id string, errNo string) error {
  return s.True(!bson.IsObjectIdHex(id), errNo)
}

// 判断字符串是否为空
func (s *Stack) StringIsEmpty(str string, errNo string) error {
  return s.True(str == "", errNo)
}

// 判断int是否为0
func (s *Stack) IntIsZero(val int, errNo string) error {
  return s.True(val == 0, errNo)
}

// 判断length是否小于
func (s *Stack) LenLessThen(str string, length int, errNo string) error {
  return s.True(len([]rune(str)) < length, errNo)
}

// 判断length是否大于
func (s *Stack) LenMoreThen(str string, length int, errNo string) error {
  return s.True(len([]rune(str)) > length, errNo)
}

// 判断时间是否早于
func (s *Stack) TimeEarlierThen(t time.Time, t2 time.Time, errNo string) error {
  return s.True(t.Before(t2), errNo)
}

// 判断时间是否晚于
func (s *Stack) TimeLaterThen(t time.Time, t2 time.Time, errNo string) error {
  return s.True(t.After(t2), errNo)
}

// 添加一个错误进栈
func (s *Stack) Error(errNo string) error {
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
