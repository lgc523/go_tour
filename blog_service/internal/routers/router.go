package routers

import (
	"github.com/gin-gonic/gin"
	global "github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/middleware"
	v1 "github.com/go-tour/blog_service/internal/routers/api/v1"
	"net/http"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.Use(middleware.Translations())

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := NewUpload()

	//file api
	e.POST("/upload/file", upload.UploadFile)
	e.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiV1 := e.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}
	return e
}
