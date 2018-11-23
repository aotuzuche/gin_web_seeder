# goweb
Simple golang restful-webservices with gin and mgo.

#### conf
配置文件，现在默认放了个apollo配置中心的玩意

----
#### controller
**控制器**

路由连接到控制器，所以正常情况下，路由数量和控制器里的方法数量是相同的。这里的方法所做的事情是：

1. 获取请求参数
2. 调用service里的操作相关数据库的方法
3. 根据service的方法返回的结果和error信息响应请求

----
#### ctx
**对gin的context的封装**

由于gin对获取参数和响应请求这些处理封装的不够，所以这里对gin的context进行一层封装，扩展一些方法。

##### getter.go: 获取参数，ctx.Get开头的方法几乎都是获取参数用的

分3类，GetParam类的，GetQuery类的，GetRaw类的，其中有类型转换，默认值这些封装。另外可自行修改或添加


##### context.go: 结构化响应内容，有ctx.Success、ctx.Error和ctx.Response

- ctx.Success

该方法需要传递一个gin.H类型的返回数据


如果是`ctx.Success(gin.H{ "a": 1 })`，则前端会收到`{ msg: 'ok', code: 0, data: { a: 1 } }`

如果是`ctx.Success(gin.H{ "data": ???? })`，则前端会解析data里的数据，这么做的目地是`data`可以直接放一个`struct`之类的数据，gin会自己处理

另外，可以直接调用`ctx.Success(nil)`，前端会由到状态码为204的返回，表示成功但没有返回内容

- ctx.Error

该方法一般传递errgo的错误码，会找到相应的errgo的错误信息并返回，也可传递error类型的，但如果传递error类型的会使用errgo的默认错误

- ctx.Response

这个方法是Success和Error的结合，接收2个参数，第一个为ctx.Error的参数，第二个为ctx.Success的参数，如果第一个参数为nil，则返回成功，否则将走错误处理的流程

----
#### db
数据库的连接、关闭操作所封装的方法

数据库一般会做如下处理：

1. 在启动项目时连接数据库，在停止项目时关闭数据库连接
2. 在请求发生时获取一个数据库操作的句柄

所以这些操作会放在这里封装成相应的方法，目前封装了redis和mongodb的，mysql的暂无

----
#### errgo

错误处理库，这里有3个文件

##### errgo.go: 维护一个error stack，这个文件一般不需要动

##### types.go: 各种错误、错误提示、http状态码的维护

自行添加时，先定义一个常量，例：`ErrSomeThingEmpty = "100001"`，然后在其下方的`Error`中添加`ErrSomeThingEmpty: { Message: "什么东西不能为空" }`。注意：`Status`可为空，默认200。`Code`不需要写，key就是其值。

然后在控制器中调用`ctx.Error(errgo.ErrSomeThingEmpty)`，前端返回`{ code: '100001', msg: '什么东西不能为空' }`并状态码为200

##### checker.go: 有验证各类错误的方法，例字符串为空、数字为0、时间大于小于某个节点、数字大于小于某个数、length大于小于某个数等等，不够的话可自行扩展。

使用方式，例：

```
str := ""
i := 300

ctx.Errgo.StringIsEmpty(str, errgo.ErrXXXXXX)
ctx.Errgo.IntIsZero(i, errgo.ErrXXXXXX)

// 捕获错误
if err := ctx.Errgo.PopError(); err != nil {
  return err
}
```
一般建议把错误验证放在service中做，这样容易复用service的方法。

----
#### middleware
gin的中间件

----
#### model
表的数据结构，建结构体，拼接结构体的数据等操作。

----
#### plugins
插件，作为就是在ctx上挂载其他常用的方法，比如在这里把mysql的数据库句柄作为插件挂载，就可以在使用时: `ctx.Sql.xxxx`这样操作，并且插件有两个生命周期钩子，分别为。

`CreatePlugins`: 他在请求发生时触发，需要返回一个`Plugins`类型，在这里比如可以做从数据库连接池中取一个句柄这类操作，并挂在`Plugins`中

`DestroyPlugins`: 在请求结束时触发，会把`CreatePlugins`方法中返回的`Plugins`再作为参数传进来，可以在这里做一个收尾工作，比如mongodb和redis都需要做关闭处理。

----
#### router
路由

----
#### service
操作数据库的方法

----
#### util
工具类