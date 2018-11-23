package errgo

import (
  "time"

  "gopkg.in/mgo.v2/bson"
)

// 判断bool返回值
func (s *Stack) True(bool bool, errNo string) error {
  if bool {
    return s.PushError(errNo)
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
