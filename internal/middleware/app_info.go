package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("app_name", "test-service")
		ctx.Set("app_version", "1.0.0")
		ctx.Next()
	}
}
