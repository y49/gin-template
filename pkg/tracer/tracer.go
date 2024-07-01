package tracer

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func NewJaegerTracer(serviceName, collectorEndpoint string) (trace.Tracer, func(), error) {
	// Configure OTLP HTTP exporter
	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(collectorEndpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create OTLP HTTP exporter: %v", err)
	}

	// Create and configure TracerProvider
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)

	// Function to close the tracer
	closer := func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to stop the tracer provider: %v", err)
		}
	}

	tracer := otel.Tracer(serviceName)
	return tracer, closer, nil
}
