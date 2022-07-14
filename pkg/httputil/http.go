package httputil

import (
	"context"
	"net/http"
	"net/url"

	"github.com/tikivn/ops-delivery-kit/pkg/opentelemetry"
	"go.opentelemetry.io/otel/trace"
)

type HttpClient struct {
	client *http.Client
}

func NewJaegerHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{client: client}
}

func (c *HttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	ctx, cleanup := opentelemetry.StartHttpSpan(req.Context(), req,
		opentelemetry.WithKind(trace.SpanKindClient),
		opentelemetry.WithCaller("http_client"),
	)
	defer cleanup(err)

	return c.client.Do(req.WithContext(ctx))
}

func (c *HttpClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	_, cleanup := opentelemetry.StartSpan(context.Background(), url,
		opentelemetry.WithKind(trace.SpanKindClient),
		opentelemetry.WithCaller("http_client"),
		opentelemetry.WithAttributes(map[string]string{
			"method": http.MethodPost,
		}),
	)
	defer cleanup(err)

	return c.client.PostForm(url, data)
}
