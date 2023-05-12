package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(

		cors.Config{

			AllowOrigins:     []string{"*"},                                                             // 等同于允许所有域名 #AllowAllOrigins:  true
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                       // 允许的请求方法
			AllowHeaders:     []string{"*", "Authorization"},                                            // 允许的请求头
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"}, // 允许的响应头
			AllowCredentials: true,                                                                      // 允许携带cookie
			MaxAge:           12 * time.Hour,                                                            // 预检请求的有效期
		},
	)

}
