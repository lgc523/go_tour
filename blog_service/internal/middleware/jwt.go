package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token    string
			respCode = errcode.Success
		)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			respCode = errcode.InvalidParam
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					respCode = errcode.AuthTimeout
				default:
					respCode = errcode.AuthFail
				}
			}
		}

		if respCode != errcode.Success {
			resp := app.NewResp(c)
			global.Logger.ErrorF("jwt check error:%s", respCode.Msg())
			resp.ErrResp(respCode)
			c.Abort()
			return
		}
		c.Next()
	}
}
