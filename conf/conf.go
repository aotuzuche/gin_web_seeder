package conf

import (
  "github.com/zouyx/agollo"
)

const (
  UseMongodb = true  // 使用mongodb
  UseRedis   = true  // 使用redis
  UseMySql   = false // 使用mysql (未实现)
)

// 获取mongodb数据库地址
func GetMongodbUrl() string {
  // return agollo.GetStringValue("mongodbUrl", "mongodb://localhost:27017/test")
  return "mongodb://localhost:27017/test"
}

// 获取redis数据库地址
func GetRedisdbUrl() string {
  return "localhost:6379"
}

// apollo配置中心的配置
func GetApolloConfig() (*agollo.AppConfig, error) {
  return &agollo.AppConfig{
    AppId:         "10200002",
    Cluster:       "dev",
    Ip:            "http://47.96.147.70:8080/",
    NamespaceName: "application",
  }, nil
}

// jwt密钥
var JwtSecret = []byte("test!@#")
