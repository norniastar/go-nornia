package controllers

import (
	"github.com/gin-gonic/gin"
	"go-nornia/service"
	"go-nornia/utils/base"
)

const (
	LoginValidNewUserErr = 10001 // 验证新用户错误
	LoginSendCodeErr     = 10002 // 发送验证码错误
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

type IsNewUserForm struct {
	Tel string `form:"tel" binding:"required"`
}

// IsNewUser 判断用户是否注册
func (c *LoginController) IsNewUser(ctx *gin.Context) { // 是否注册
	m := new(IsNewUserForm)
	if err := ctx.ShouldBind(&m); err != nil {
		base.ErrorParam(ctx, err)
		return
	}

	ret, err := service.NewUserService().VerifyUser(m.Tel)
	if err != nil {
		base.FailResponse(ctx, LoginValidNewUserErr, err)
		return
	}

	base.SuccessResponse(ctx, ret)
	return
}
