package limiter

import (
    "github.com/juju/ratelimit"
    "github.com/gin-gonic/gin"
    "time"
)

type LimiterInterface interface {
    Key(c *gin.Context) string
    GetBucket(key string) (*ratelimit.Bucket, bool)
    AddBuckets(rules ...LimiterBucketRule) LimiterInterface
}

type Limiter struct {
    limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
    Key          string
    FillInterval time.Duration
    Capacity     int64
    Quantum      int64
}