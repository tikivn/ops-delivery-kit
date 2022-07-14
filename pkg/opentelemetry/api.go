package opentelemetry

import (
	"context"
	"net/http"
)

func JaegerWrapperBefore(ctx context.Context, r *http.Request) context.Context {
	/*
		tracer := TracerFromContext(ctx)

		spanAttributes := []attribute.KeyValue{
			{Key: "path", Value: attribute.StringValue(r.URL.Path)},
			{Key: "param", Value: attribute.StringValue(r.URL.Query().Encode())},
		}

		ctxTrace, _ := tracer.Start(ctx, r.URL.String(),
			trace.WithNewRoot(),
			trace.WithAttributes(spanAttributes...),
		)

		return ctxTrace
	*/
	return ctx
}

func JaegerWrapperAfter(ctx context.Context, w http.ResponseWriter) context.Context {
	// span := trace.SpanFromContext(ctx)
	// if span.IsRecording() {
	// 	span.End()
	// }

	return ctx
}
