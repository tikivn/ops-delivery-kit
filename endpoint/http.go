package endpoint

import (
	"context"
	"net/http"

	"github.com/getsentry/raven-go"
)

func SentryMiddleware(h http.Handler) http.Handler {
	withRecovery := raven.Recoverer(h)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		withRecovery.ServeHTTP(w, NewRequest(r))
	})
}

func NewRequest(r *http.Request) *http.Request {
	return r.WithContext(WithRavenHttp(r.Context(), r))
}

type ravenHttpKeyType struct{}

var (
	ravenHttpKey ravenHttpKeyType
)

func WithRavenHttp(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, ravenHttpKey, raven.NewHttp(r))
}

func HttpFromContext(ctx context.Context) *raven.Http {
	val := ctx.Value(ravenHttpKey)
	if h, ok := val.(*raven.Http); ok {
		return h
	}
	return nil
}

// Re-write http-code 301 && 302 to 307 && 308
type ResponseWriter struct {
	http.ResponseWriter
	status               int
	retainRedirectBodies bool
	wroteHeader          bool
}

func NewResponseWriter(w http.ResponseWriter, retainRedirectBodies bool) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w, retainRedirectBodies: retainRedirectBodies}
}

func (w *ResponseWriter) Status() int {
	return w.status
}

func (w *ResponseWriter) Write(p []byte) (n int, err error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(p)
}

func (w *ResponseWriter) WriteHeader(code int) {
	if w.retainRedirectBodies {
		switch code {
		case http.StatusFound:
			code = 308
		case http.StatusMovedPermanently:
			code = 307
		}
	}

	w.ResponseWriter.WriteHeader(code)
	if w.wroteHeader {
		return
	}
	w.status = code
	w.wroteHeader = true
}

func RetainBodyOnRedirect(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wr := NewResponseWriter(w, true)
		h.ServeHTTP(wr, r)
	}

	return http.HandlerFunc(fn)
}
