package controller

import (
	"errors"
	"fmt"
	"gin_demo/dao"
	"gin_demo/dto"
	"gin_demo/middleware"
	"gin_demo/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// Register godoc
// @Summary      用户注册
// @Description  用户注册接口
// @Tags         用户管理
// @ID           /user/register
// @Accept       json
// @Produce      json
// @Param        body body      dto.RegisterInput                          true  "body"
// @Success      200   {object}  middleware.Response{data=string}  "success"
// @Router       /user/register [post]
func (u *UserController) Register(c *gin.Context) {
	params := &dto.RegisterInput{}
	if err := dto.BindValidParams(c, params); err != nil {
		middleware.ResponseError(c, err)
		return
	}
	userDao := dao.User{}
	user, err := (&userDao).FindOne(map[string]interface{}{
		"account": params.Account,
	})
	if err != nil {
		middleware.ResponseError(c, err)

	} else if user.Password != "" {
		errMsg := fmt.Sprintf("账号:%s已存在", params.Account)
		middleware.ResponseError(c, errors.New(errMsg))
	} else {
		(&userDao).Create(&dao.User{
			Name:     params.Name,
			Account:  params.Account,
			Password: utils.EncodeByMd5(params.Password),
			Age:      params.Age,
			Birth:    params.Birth,
			Sex:      params.Sex,
		})
		middleware.ResponseSuccess(c, "创建成功")
	}
}

// Login godoc
// @Summary      用户登录
// @Description  用户登录接口
// @Tags         用户管理
// @ID           /user/login
// @Accept       json
// @Produce      json
// @Param        body body      dto.LoginInput                          true  "body"
// @Success      200   {object}  middleware.Response{data=dto.LoginOutput}  "success"
// @Router       /user/login [post]
func (u *UserController) Login(c *gin.Context) {
	params := &dto.LoginInput{}
	_ = c.ShouldBindJSON(params)
	userDao := dao.User{}
	user, err := (&userDao).FindOne(map[string]interface{}{
		"account":params.Account,
		"password":utils.EncodeByMd5(params.Password),
	})
	if err != nil {
		middleware.ResponseError(c, err)
	} else if user.Id == 0 {
		middleware.ResponseError(c,errors.New("账号或密码错误"))
	} else {
		token, err := utils.GenToken(user.Account)
		if err != nil {
			middleware.ResponseError(c, err)
		} else {
			middleware.ResponseSuccess(c, map[string]interface{}{
				"token": token,
			})
		}
	}
}
// GetUserInfo godoc
// @Summary      用户获取个人信息
// @Description  用户获取个人信息接口
// @Tags         用户管理
// @ID           /user/getUserInfo
// @Param        authorization header string                                true  "header"
// @Accept       json
// @Success      200   {object}  middleware.Response{data=dto.GetUserInfoOutput}  "success"
// @Router       /user/getUserInfo [get]
func (u *UserController) GetUserInfo(c *gin.Context) {
	account, _ := c.Get("account")

	daoUser := dao.User{}
	user, err := (&daoUser).FindOne(map[string]interface{}{
		"account":account.(string),
	})
	if err != nil {
		middleware.ResponseError(c, err)
	} else {
		middleware.ResponseSuccess(c, map[string]interface{}{
			"id":      user.Id,
			"name":    user.Name,
			"account": user.Account,
			"age":     user.Age,
			"birth":   user.Birth,
			"sex":     user.Sex,
			"avatar_url": user.AvatarUrl,
		})
	}
}
