package middleware

import (
	"errors"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			ResponseError(c,errors.New("请求头中auth为空"))
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			ResponseError(c, errors.New("请求头中auth格式有误"))
			c.Abort()
			return
		}
		data, err := utils.ParseToken(authHeader)
		if err != nil {
			ResponseError(c,errors.New("无效的Token"))
			c.Abort()
			return
		}
		//m := mc.(jwt.MapClaims)
		//// 将当前请求的username信息保存到请求的上下文c上
		c.Set("account", data.Account)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}


