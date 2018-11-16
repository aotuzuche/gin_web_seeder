package util

import "github.com/gin-gonic/gin"

func Exist(k string, keys ... string) bool {
  for _, i := range keys {
    if i == k {
      return true
    }
  }
  return false
}

// 删除ginH的部分数据
func IgnoreData(m gin.H, keys ... string) {
  for _, v := range keys {
    if _, ok := m[v]; ok {
      delete(m, v)
    }
  }
}

// 保留ginH的部分数据
func RetainData(m gin.H, keys ... string) {
  for k := range m {
    if !Exist(k, keys...) {
      delete(m, k)
    }
  }
}
