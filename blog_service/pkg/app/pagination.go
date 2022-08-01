package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/pkg/convert"
)

func GetPage(c *gin.Context) int {
	pageNo := convert.StrTo(c.Query("pageNo")).MustInt()
	if pageNo <= 0 {
		return 1
	}
	return pageNo
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("pageSize")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(pageNo, pageSize int) int {
	result := 0
	if pageNo > 0 {
		result = (pageNo - 1) * pageSize
	}
	return result
}
