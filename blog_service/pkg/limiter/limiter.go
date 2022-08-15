package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"sync"
	"time"
)

type LimitService interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimitBucketRule) LimitService
}

type Limiter struct {
	sync.Mutex
	limiterBuckets map[string]*ratelimit.Bucket
	//map[string]*ratelimit.Bucket
	//lock           sync.Mutex
}

type LimitBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quota        int64 //set token count every fillInterval
}
