package consumer

import (
	"context"
)

type Tracer interface {
	Transaction(ctx context.Context, name string) (context.Context, func(err error))
	Segment(ctx context.Context, name string) (context.Context, func(err error))
}

type noopTracing struct{}

func (_ noopTracing) Transaction(ctx context.Context, name string) (context.Context, func(_ error)) {
	return ctx, func(_ error) {}
}

func (_ noopTracing) Segment(ctx context.Context, name string) (context.Context, func(_ error)) {
	return ctx, func(_ error) {}
}

type newrelicProcessor struct {
	Processor
	tracer Tracer
}

func NewProcessor(p Processor, tracer Tracer) Processor {
	return &newrelicProcessor{
		Processor: p,
		tracer:    tracer,
	}
}

func (p *newrelicProcessor) Decode(ctx context.Context, data []byte) (msg interface{}, err error) {
	ctx, done := p.tracer.Segment(ctx, "decode")
	defer func() {
		done(err)
	}()
	return p.Processor.Decode(ctx, data)
}

func (p *newrelicProcessor) Process(ctx context.Context, msg interface{}) (err error) {
	ctx, done := p.tracer.Segment(ctx, "process")
	defer func() {
		done(err)
	}()
	return p.Processor.Process(ctx, msg)
}
