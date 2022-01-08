package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(ctx *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rulers ...LimiterBucketRuler) LimiterIface
}

type Limiter struct {
	LimiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRuler struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
