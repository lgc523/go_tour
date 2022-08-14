package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/pkg/app"
)

func AppInfo(c *gin.Context) {
	resp := app.NewResp(c)
	resp.Success(gin.H{
		"appName":    c.MustGet("appName"),
		"appVersion": c.MustGet("appVersion"),
	})
}
