package ctx

import (
  "strconv"
  "strings"
  "time"

  "github.com/tidwall/gjson"
  "gopkg.in/mgo.v2/bson"
)


// 从body的raw中注入到一个结构体
func (c *New) InjectRaw(value interface{}) {
  bson.UnmarshalJSON(c.RawData, value)
}

// 从body中取string类型的值
func (c *New) GetRaw(key string) (string, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return strings.TrimSpace(res.Str), res.Exists()
}

// 从body中取string类型的值，没有则使用默认值
func (c *New) GetRawDefault(key string, def string) string {
  val, ok := c.GetRaw(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取array类型的值
func (c *New) GetRawArray(key string) ([]gjson.Result, bool) {
  res := gjson.GetBytes(c.RawData, key)
  arr := res.Array()
  return arr, res.Exists()
}

// 从body中取array类型的值，没有则使用默认值
func (c *New) GetRawArrayDefault(key string, def []gjson.Result) []gjson.Result {
  val, ok := c.GetRawArray(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取time类型的值
func (c *New) GetRawTime(key string) (time.Time, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return res.Time(), res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawTimeDefault(key string, def time.Time) time.Time {
  val, ok := c.GetRawTime(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取int类型的值
func (c *New) GetRawInt(key string) (int, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return int(res.Int()), res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawIntDefault(key string, def int) int {
  val, ok := c.GetRawInt(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取bool类型的值
func (c *New) GetRawBool(key string) (bool, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return res.Bool(), res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawBoolDefault(key string, def bool) bool {
  val, ok := c.GetRawBool(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取gjson类型的值
func (c *New) GetRawGJSON(key string) (gjson.Result, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return res, res.Exists()
}

// 从params中取string类型的值
func (c *New) GetParam(key string) (string, bool) {
  res, ok := c.Ctx.Params.Get(key)
  return res, ok
}

// 从params中取string类型的值，没有则使用默认值
func (c *New) GetParamDefault(key string, def string) string {
  val, ok := c.GetParam(key)
  if !ok {
    return def
  }
  return val
}

// 从params中取int类型的值
func (c *New) GetParamInt(key string) (int, bool) {
  res, ok := c.Ctx.Params.Get(key)
  intRes, _ := strconv.Atoi(res)
  return intRes, ok
}

// 从params中取int类型的值，没有则使用默认值
func (c *New) GetParamIntDefault(key string, def int) int {
  val, ok := c.GetParamInt(key)
  if !ok {
    return def
  }
  return val
}

// 从params中取bool类型的值
func (c *New) GetParamBool(key string) (bool, bool) {
  res, ok := c.Ctx.Params.Get(key)
  return res == "true", ok
}

// 从params中取bool类型的值，没有则使用默认值
func (c *New) GetParamBoolDefault(key string, def bool) bool {
  val, ok := c.GetParamBool(key)
  if !ok {
    return def
  }
  return val
}

// 从query中取string类型的值
func (c *New) GetQuery(key string) (string, bool) {
  res, ok := c.Ctx.GetQuery(key)
  return res, ok
}

// 从query中取string类型的值，没有则使用默认值
func (c *New) GetQueryDefault(key string, def string) string {
  val, ok := c.GetQuery(key)
  if !ok {
    return def
  }
  return val
}

// 从query中取int类型的值
func (c *New) GetQueryInt(key string) (int, bool) {
  res, ok := c.Ctx.GetQuery(key)
  if !ok {
    return 0, false
  }
  intRes, err := strconv.Atoi(res)
  if err != nil {
    return 0, false
  }
  return intRes, true
}

// 从query中取int类型的值，没有则使用默认值
func (c *New) GetQueryIntDefault(key string, def int) int {
  val, ok := c.GetQueryInt(key)
  if !ok {
    return def
  }
  return val
}

// 从query中取bool类型的值
func (c *New) GetQueryBool(key string) (bool, bool) {
  res, ok := c.Ctx.GetQuery(key)
  if !ok {
    return false, false
  }
  return res == "true", true
}

// 从query中取bool类型的值，没有则使用默认值
func (c *New) GetQueryBoolDefault(key string, def bool) bool {
  val, ok := c.GetQueryBool(key)
  if !ok {
    return def
  }
  return val
}

// 从ctx中取字符串类型的值
func (c *New) Get(key string) (string, bool) {
  res, ok := c.Ctx.Get(key)
  if !ok {
    return "", false
  }
  return res.(string), true
}

// 从ctx中取int类型的值
func (c *New) GetInt(key string) (int, bool) {
  res, ok := c.Ctx.Get(key)
  if !ok {
    return 0, false
  }
  return res.(int), true
}

// 从ctx中取bool类型的值
func (c *New) GetBool(key string) (bool, bool) {
  res, ok := c.Ctx.Get(key)
  if !ok {
    return false, false
  }
  return res.(bool), true
}
