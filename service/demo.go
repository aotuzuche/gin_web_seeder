package service

import (
  "errors"
  "web/ctx"
  "web/errgo"
  "web/model"
  "web/util"

  "gopkg.in/mgo.v2/bson"
)

// demo server
func Insert(ctx *ctx.New, data model.DemoModel) error {
  // 验证字段
  ctx.Errgo.StringIsEmpty(data.Foo, errgo.ErrForbidden)
  ctx.Errgo.IntLessThen(data.Bar, 0, errgo.ErrSkipRange)
  ctx.Errgo.IntLessThen(data.Bar, 1, errgo.ErrLimitRange)
  ctx.Errgo.IntMoreThen(data.Bar, 50, errgo.ErrLimitRange)

  // 字段有误则返回
  if err := ctx.Errgo.PopError(); err != nil {
    return err
  }

  // 存
  err := ctx.MgoDB.C(model.DemoCollection).Insert(data)

  // 返回结果
  if err != nil {
    return errors.New(errgo.ErrServerError)
  }
  return nil
}

func Update(ctx *ctx.New, id string, data bson.M) error {
  // 限制更新的字段及类型
  util.Only(
    data,
    util.Keys{
      "foo": util.TypeString,
    },
  )

  // 验证id
  ctx.Errgo.StringNotObjectId(id, errgo.ErrIdError)

  if err := ctx.Errgo.PopError(); err != nil {
    return err
  }

  // 更新
  err := ctx.MgoDB.C(model.DemoCollection).UpdateId(bson.ObjectIdHex(id), bson.M{
    "$set": data,
  })

  // 返回结果
  if err != nil {
    return errors.New(errgo.ErrServerError)
  }

  return nil
}
