package limiter

import (
    "github.com/gin-gonic/gin"
    "github.com/juju/ratelimit"
    "strings"
)

type MethodLimiter struct {
    *Limiter
}

func (ml MethodLimiter) Key(c *gin.Context) string{
    uri := c.Request.RequestURI
    index := strings.Index(uri, "?")
    if index == -1 {
        return uri
    }

    return uri[:index]
}

func (ml MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
    bucket, ok := ml.limiterBuckets[key]
    return bucket, ok
}

func (ml MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
    for _, rule := range rules {
        if _, ok := ml.limiterBuckets[rule.Key]; !ok {
            ml.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
        }
    }

    return ml
}

func NewMethodLimiter() LimiterInterface {
    return MethodLimiter{
        Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
    }
}

