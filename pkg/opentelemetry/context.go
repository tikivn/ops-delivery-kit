package opentelemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type tracerKeyStruct struct{}

var tracerKey tracerKeyStruct

func TracerFromContext(ctx context.Context) trace.Tracer {
	val := ctx.Value(tracerKey)
	if tracer, ok := val.(trace.Tracer); ok {
		return tracer
	}

	return otel.GetTracerProvider().Tracer(TRACER_DEFAULT_OPERATIONS)
}

func ContextWithTracer(ctx context.Context, tracer trace.Tracer) context.Context {
	return context.WithValue(ctx, tracerKey, tracer)
}
