package middleware

import "github.com/gin-gonic/gin"

type Response struct {
	Code int8 `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, err error) {
	resp := &Response {
		Code: -1,
		Message: err.Error(),
		Data:"",
	}
	c.JSON(200, resp)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{
		Code: 0,
		Message: "请求成功",
		Data: data,
	}
	c.JSON(200,resp)
}
