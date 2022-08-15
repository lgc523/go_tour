package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimitService {
	return &MethodLimiter{
		Limiter: &Limiter{
			limiterBuckets: make(map[string]*ratelimit.Bucket),
			//lock:           sync.Mutex{},
		},
	}
}

func (l *MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (l *MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	l.Lock()
	bucket, ok := l.limiterBuckets[key]
	defer l.Unlock()
	return bucket, ok
}

func (l *MethodLimiter) AddBuckets(rules ...LimitBucketRule) LimitService {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			l.Lock()
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quota)
			l.Unlock()
		}
	}
	return l
}
