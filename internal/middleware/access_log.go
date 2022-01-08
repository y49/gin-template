package middleware

import (
	"bytes"
	"gin-template/global"
	"gin-template/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = bodyWriter

		beginTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method:%S,status_code:%d," + "begin_time:%d,end_time:%d"
		global.Logger.WithFields(fields).Infof(ctx, s, ctx.Request.Method, bodyWriter.Status(), beginTime, endTime)
	}
}
