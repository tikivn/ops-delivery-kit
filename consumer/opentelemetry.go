package consumer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type opentelemetryTracer struct {
	tracer trace.Tracer
}

func OpenTelemetryTracer(tracer trace.Tracer) Tracer {
	return &opentelemetryTracer{
		tracer: tracer,
	}
}

func (o *opentelemetryTracer) Transaction(ctx context.Context, name string) (context.Context, func(err error)) {
	ctx, span := o.tracer.Start(ctx, name,
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)

	return ctx, func(err error) {
		span.RecordError(err)
		span.End()
	}
}

func (o *opentelemetryTracer) Segment(ctx context.Context, name string) (context.Context, func(err error)) {
	traceCtx, span := o.tracer.Start(ctx, name,
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithLinks(trace.LinkFromContext(ctx)),
	)

	return traceCtx, func(err error) {
		span.RecordError(err)
		span.End()
	}
}
