package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	Name string `json:"name" from:"name" validate:"required" comment:"用户名"` //用户名
	Account string `json:"account" from:"account" validate:"required" comment:"账号"` //账号
	Password string `json:"password" from:"password" validate:"min=6,max=22" comment:"密码"` //密码
	Age uint8 `json:"age" from:"age" validate:"gt=0" comment:"年龄"`//年龄
	Birth string `json:"birth" from:"birth" validate:"required" comment:"出生日期"`//出生日期
	Sex int `json:"sex" from:"sex" validate:"required" comment:"性别 1：男 2：女"`//性别 1：男 2：女
}

type LoginInput struct {
	Account string `json:"account" from:"account" validate:"min=6,max=22,required" comment:"账号"`//账号
	Password string `json:"password" from:"password" validate:"required" comment:"密码"`//密码
}

type LoginOutput struct {
	Token string `json:"token"` //用户token
}

type GetUserInfoOutput struct {
	Id string `json:"id"` //用户id
	Name string `json:"name"` //用户名
	Account string `json:"account"` //账号
	Age uint8 `json:"age"` //年龄
	Birth string `json:"birth"` //出生日期
	Sex string `json:"sex"`//性别 1：男 2：女
	AvatarUrl string `json:"avatar_url"` //头像
}

func BindValidParams(c *gin.Context, params interface{}) error  {
	err := c.ShouldBindJSON(params)
	if err != nil {
		return err
	}

	validate := validator.New()
	return validate.Struct(params)
}
