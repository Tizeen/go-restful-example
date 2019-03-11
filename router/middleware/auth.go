package middleware

import (
	"github.com/Tizeen/go-restful-example/handler"
	"github.com/Tizeen/go-restful-example/pkg/errno"
	"github.com/Tizeen/go-restful-example/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			// 停止当前handler。认证失败直接退出
			c.Abort()
			return
		}

		c.Next()
	}
}
