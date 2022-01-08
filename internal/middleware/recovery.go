package middleware

import (
	"fmt"
	"gin-template/global"
	"gin-template/pkg/app"
	"gin-template/pkg/email"
	"gin-template/pkg/errcode"
	"time"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover error: %v"
				global.Logger.WithCallersFrames().Errorf(ctx, s, err)

				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出时间：%d", time.Now().Unix()),
					fmt.Sprintf("信息: %v", err),
				)
				if err != nil {
					global.Logger.Panic(ctx, "mail.SendMail error: %v", err)
				}

				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
