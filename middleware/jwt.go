package middleware

import (
  "regexp"
  "web/conf"
  "web/ctx"
  "web/errgo"
  "workerbook/context"

  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
)

// check up json web token
func Jwt(c *gin.Context) {
  headerAuth, headerJwt := c.Request.Header.Get("authorization"), ""

  jwtReg := regexp.MustCompile(`^Bearer\s\S+$`)

  if jwtReg.MatchString(headerAuth) {
    headerJwt = headerAuth[len("Bearer "):]

    token, _ := jwt.Parse(headerJwt, func(t *jwt.Token) (interface{}, error) {
      return conf.JwtSecret, nil
    })

    if token.Valid {
      c.Next()
    } else {
      bctx := ctx.NewBaseCtx(c)
      bctx.Error(errgo.ErrNeedLogin)
      c.Abort()
    }
  } else {
    bctx := context.NewBaseCtx(c)
    bctx.Error(errgo.ErrNeedLogin)
    c.Abort()
  }
}
