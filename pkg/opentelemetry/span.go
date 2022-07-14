package opentelemetry

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type spanFunctionConfig struct {
	// functionName string
	param      []interface{}
	identify   string
	caller     string
	kind       trace.SpanKind
	attributes map[string]string
}

type SpanFunctionOption interface {
	applyConfig(config spanFunctionConfig) spanFunctionConfig
}

type spanFunction func(config spanFunctionConfig) spanFunctionConfig

func (fn spanFunction) applyConfig(cfg spanFunctionConfig) spanFunctionConfig {
	return fn(cfg)
}

func WithIdentifier(identify string) SpanFunctionOption {
	return spanFunction(func(cfg spanFunctionConfig) spanFunctionConfig {
		cfg.identify = identify
		return cfg
	})
}

func WithParameter(params []interface{}) SpanFunctionOption {
	return spanFunction(func(cfg spanFunctionConfig) spanFunctionConfig {
		cfg.param = params
		return cfg
	})
}

func WithCaller(caller string) SpanFunctionOption {
	return spanFunction(func(cfg spanFunctionConfig) spanFunctionConfig {
		cfg.caller = caller
		return cfg
	})
}

func WithKind(kind trace.SpanKind) SpanFunctionOption {
	return spanFunction(func(cfg spanFunctionConfig) spanFunctionConfig {
		cfg.kind = kind
		return cfg
	})
}

func WithAttributes(attr map[string]string) SpanFunctionOption {
	return spanFunction(func(cfg spanFunctionConfig) spanFunctionConfig {
		if cfg.attributes == nil {
			cfg.attributes = make(map[string]string)
		}

		for k, v := range attr {
			cfg.attributes[k] = v
		}
		return cfg
	})
}

func newSpanFuncConfig(options []SpanFunctionOption) spanFunctionConfig {
	var cfg spanFunctionConfig
	for _, o := range options {
		cfg = o.applyConfig(cfg)
	}
	return cfg
}

func StartSpan(ctx context.Context, function string, options ...SpanFunctionOption) (context.Context, func(error)) {
	cfg := newSpanFuncConfig(options)

	spanOps := spanStartOption(cfg)
	spanOps = append(spanOps, trace.WithLinks(trace.LinkFromContext(ctx)))

	ctxTrace, span := TracerFromContext(ctx).Start(ctx, function,
		spanOps...,
	)

	return ctxTrace, func(err error) {
		span.RecordError(err, trace.WithStackTrace(err != nil))
		span.End()
	}
}

func StartHttpSpan(ctx context.Context, request *http.Request, options ...SpanFunctionOption) (context.Context, func(error)) {
	cfg := newSpanFuncConfig(options)

	spanOps := spanStartOption(cfg, trace.WithAttributes(semconv.HTTPClientAttributesFromHTTPRequest(request)...))
	spanOps = append(spanOps, trace.WithLinks(trace.LinkFromContext(ctx)))

	ctxTrace, span := TracerFromContext(ctx).Start(ctx, request.URL.Path,
		spanOps...,
	)

	return ctxTrace, func(err error) {
		span.RecordError(err, trace.WithStackTrace(err != nil))
		span.End()
	}
}

func attributeFromConfig(cfg spanFunctionConfig) []attribute.KeyValue {
	attr := make([]attribute.KeyValue, 0)

	if cfg.identify != "" {
		attr = append(attr, attribute.KeyValue{
			Key:   "identify",
			Value: attribute.StringValue(cfg.identify),
		})
	}

	if cfg.param != nil {
		paramSlice := make([]string, 0, len(cfg.param))
		for _, v := range cfg.param {
			paramSlice = append(paramSlice, fmt.Sprintf("%+v", v))
		}

		attr = append(attr, attribute.KeyValue{
			Key:   "parameter",
			Value: attribute.StringSliceValue(paramSlice),
		})
	}

	if cfg.caller != "" {
		attr = append(attr, attribute.KeyValue{
			Key:   "caller",
			Value: attribute.StringValue(cfg.caller),
		})
	}

	if cfg.attributes != nil {
		for k, v := range cfg.attributes {
			attr = append(attr, attribute.KeyValue{
				Key:   attribute.Key(k),
				Value: attribute.StringValue(v),
			})
		}
	}
	return attr
}

func spanStartOption(cfg spanFunctionConfig, additionalStartOptions ...trace.SpanStartOption) []trace.SpanStartOption {
	options := make([]trace.SpanStartOption, 0)
	options = append(options, trace.WithSpanKind(cfg.kind))

	attr := attributeFromConfig(cfg)
	options = append(options, trace.WithAttributes(attr...))

	return append(options, additionalStartOptions...)
}
