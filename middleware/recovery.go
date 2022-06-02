package middleware

import (
	"errors"
	"fmt"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)
var config = utils.Config
func RecoveryMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))

				if config.DebugMode != "debug" {
					ResponseError(c, errors.New("内部错误"))
				} else {
					ResponseError(c,errors.New(fmt.Sprint(err)))
				}
				return
			}
		}()
		c.Next()
	}
}
