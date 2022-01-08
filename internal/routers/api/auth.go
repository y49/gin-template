package api

import (
	"gin-template/global"
	"gin-template/internal/service"
	"gin-template/pkg/app"
	"gin-template/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	params := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svg := service.New(c.Request.Context())
	err := svg.CheckAuth(&params)
	if err != nil {
		global.Logger.Errorf(c, "svg.CheckAuth: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GetJWTToken(params.Appkey, params.Appsecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GetJWTToken: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
