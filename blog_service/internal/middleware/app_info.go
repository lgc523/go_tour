package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("appName", "blog_service")
		c.Set("appVersion", "1.0.0")
		c.Next()
	}
}
