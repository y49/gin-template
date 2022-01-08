package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	l := &Limiter{LimiterBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimiter{
		Limiter: l,
	}
}

func (l MethodLimiter) Key(ctx *gin.Context) string {
	url := ctx.Request.RequestURI
	index := strings.Index(url, "?")
	if index == -1 {
		return url
	}
	return url[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.LimiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rulers ...LimiterBucketRuler) LimiterIface {
	for _, rule := range rulers {
		if _, ok := l.LimiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval,
				rule.Capacity,
				rule.Quantum,
			)
			l.LimiterBuckets[rule.Key] = bucket
		}
	}
	return l
}
