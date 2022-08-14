package routers

import (
	"github.com/gin-gonic/gin"
	global "github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/middleware"
	"github.com/go-tour/blog_service/internal/routers/api"
	v1 "github.com/go-tour/blog_service/internal/routers/api/v1"
	"github.com/go-tour/blog_service/pkg/limiter"
	"net/http"
	"time"
)

func NewRouter() *gin.Engine {
	e := gin.New()

	//if global.ServerSetting.RunMode == "debug" {
	//	e.Use(gin.Logger())
	//	e.Use(gin.Recovery())
	//} else {
	e.Use(middleware.AccessLog())
	e.Use(middleware.Recovery())
	//}
	e.Use(middleware.Tracing())

	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
		limiter.LimitBucketRule{
			Key:          "/info",
			FillInterval: time.Second,
			Capacity:     10,
			Quota:        10,
		},
	)

	e.Use(middleware.Translations())
	e.Use(middleware.AppInfo())
	e.Use(middleware.RateLimiter(methodLimiters))
	e.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	auth := api.NewAuth()

	//file api
	e.POST("/upload/file", upload.UploadFile)
	e.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	//jwt
	e.GET("/auth", auth.GetAuth)

	//info
	e.GET("/info", api.AppInfo)

	apiV1 := e.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/article", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.POST("/articles", article.List)
	}
	return e
}
