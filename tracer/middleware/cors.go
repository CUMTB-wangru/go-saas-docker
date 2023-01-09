package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			// AllowOrigins 允许访问的域名列表，*表是所有域名都也已访问，默认值[]
			AllowOrigins:     []string{"*"},
			// AllowMethods 允许客户端使用的请求方法列表，默认get/post
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			// AllowHeaders 允许客户端使用的非简单头的列表
			AllowHeaders:     []string{"*"},
			// ExposeHeaders 指示哪些头可以安全地暴露于CORS的API
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			// AllowCredentials 允许添加源
			AllowCredentials: true,
			// MaxAge 响应时间
			MaxAge:           12 * time.Hour,
		},
	)
}
