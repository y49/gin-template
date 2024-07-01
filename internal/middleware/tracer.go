package middleware

import (
	"fmt"
	"gin-template/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract parent context from incoming request
		propagator := propagation.TraceContext{}
		parentCtx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		tracer := otel.Tracer(global.AppSetting.ServiceName)
		spanName := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)

		// Start a new span
		spanCtx, span := tracer.Start(parentCtx, spanName)
		defer span.End()

		// Record some basic attributes
		span.SetAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.url", c.Request.URL.String()),
			attribute.String("http.client_ip", c.ClientIP()),
		)

		// Add trace IDs to the context
		c.Set("X-Tracer-ID", span.SpanContext().TraceID().String())
		c.Set("X-Span-ID", span.SpanContext().SpanID().String())

		// Add the span context to the request context
		c.Request = c.Request.WithContext(spanCtx)

		// Measure request processing time
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Record the status code and response time
		span.SetAttributes(
			attribute.Int("http.status_code", c.Writer.Status()),
			attribute.Float64("http.response_time_ms", float64(duration.Milliseconds())),
		)

		// Log error if status code is 5xx
		if c.Writer.Status() >= 500 {
			span.SetAttributes(attribute.String("error", "true"))
			span.SetAttributes(attribute.String("error.message", c.Errors.ByType(gin.ErrorTypePrivate).String()))
		}
	}
}
