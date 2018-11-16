package model

import (
  "web/util"

  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2/bson"
)

// 文档名
const DemoCollection = "demos"

// 文档的字段结构
type DemoModel struct {
  Id  bson.ObjectId `bson:"_id,omitempty"` // id自动生成
  Foo string        `bson:"foo"`
  Bar int           `bson:"bar"`
}

// 实现接口
func (d DemoModel) Parse() gin.H {
  return gin.H{
    "id":  d.Id,
    "foo": d.Foo,
    "len": len(d.Foo),
  }
}

// 解析并过滤
func (d DemoModel) ParseAndIgnore(ignore ... string) gin.H {
  data := d.Parse()
  if ignore != nil {
    util.IgnoreData(data, ignore...)
  }
  return data
}

// 解析并保留
func (d DemoModel) ParseAndRetain(retain ... string) gin.H {
  data := d.Parse()
  if retain != nil {
    util.RetainData(data, retain...)
  }
  return data
}
