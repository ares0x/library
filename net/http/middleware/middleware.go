package middleware

import (
    "context"
    "github.com/gin-gonic/gin"
    opentracing "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
    "github.com/uber/jaeger-client-go"
    "library/limiter"
    "time"
)


//链路追踪
func Tracing(tracer opentracing.Tracer) func(c *gin.Context) {
    return func(c *gin.Context) {
        var (
            ctx context.Context
            span opentracing.Span
        )
        spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
        if err != nil {
            span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),tracer, c.Request.URL.Path )
        } else {
            span, ctx = opentracing.StartSpanFromContextWithTracer(
                c.Request.Context(),
                tracer,
                c.Request.URL.Path,
                opentracing.ChildOf(spanCtx),
                opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
                )
        }
        defer span.Finish()

        var traceID string
        var spanID string
        var spanContext = span.Context()
        switch spanContext.(type) {
        case jaeger.SpanContext:
            jaegerContext := spanContext.(jaeger.SpanContext)
            traceID = jaegerContext.TraceID().String()
            spanID = jaegerContext.SpanID().String()
        }
        c.Set("X-Trace-ID", traceID)
        c.Set("X-Span-ID", spanID)
        c.Request = c.Request.WithContext(ctx)
        c.Next()
    }
}

//接口限流
func RateLimiter(li limiter.LimiterInterface) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := li.Key(c)
        if bucket, ok := li.GetBucket(key); ok {
            count := bucket.TakeAvailable(1)
            if count == 0{

            }
        }
    }
}

//超时控制
func ContextTimeOut(t time.Duration) func(c *gin.Context) {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(c.Request.Context(), t)
        defer cancel()

        c.Request = c.Request.WithContext(ctx)
        c.Next()
    }
}