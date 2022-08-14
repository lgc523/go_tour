package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/pkg/app"
	"github.com/go-tour/blog_service/pkg/errcode"
	"github.com/go-tour/blog_service/pkg/limiter"
)

func RateLimiter(l limiter.LimitService) gin.HandlerFunc {
	return func(c *gin.Context) {
		routeUrl := l.Key(c)
		if bucket, ok := l.GetBucket(routeUrl); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				resp := app.NewResp(c)
				resp.ErrResp(errcode.TooManyReq)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
