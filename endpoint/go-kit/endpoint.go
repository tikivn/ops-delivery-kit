package go_kit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/getsentry/raven-go"
	endpoint_kit "github.com/go-kit/kit/endpoint"
	"go.opencensus.io/trace"

	"github.com/tikivn/ops-delivery-kit/endpoint"
)

// EncodeError encode errors from business-logic
func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	res := &endpoint.BasicResponse{Status: endpoint.StatusError}
	var statusCode int
	switch e := err.(type) {
	case endpoint.ErrBadRequest:
		res.Error = map[string]interface{}{
			"code":    "VALIDATION_ERROR",
			"message": e.Error(),
		}

		statusCode = http.StatusBadRequest
	case endpoint.BizError:
		res.Error = map[string]interface{}{
			"code":    e.Code(),
			"message": e.Error(),
		}

		statusCode = http.StatusOK
	default:
		fmt.Println("Error happen:", err)
		packet := raven.NewPacket(
			err.Error(),
			raven.NewException(err, raven.GetOrNewStacktrace(err, 2, 3, nil)),
			endpoint.HttpFromContext(ctx),
		)
		raven.Capture(packet, nil)
		res.Error = endpoint.ErrInternal

		statusCode = http.StatusInternalServerError
	}

	if data, ok := err.(endpoint.ExtraData); ok {
		res.Error["data"] = data.Data()
	}

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(res)
}

func EncodeResponse(
	ctx context.Context,
	w http.ResponseWriter,
	data interface{},
) error {
	if f, ok := data.(endpoint_kit.Failer); ok {
		if e := f.Failed(); e != nil {
			EncodeError(ctx, e, w)
			return nil
		}
	}

	res := &endpoint.BasicResponse{
		Status: endpoint.StatusOk,
		Data:   data,
	}

	if p, ok := data.(endpoint.Pager); ok {
		res.Paging = endpoint.NewPaging(p.Pagination())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	span := trace.FromContext(ctx)
	if span != nil {
		spanCtx := span.SpanContext()
		w.Header().Set("x-trace-id", spanCtx.TraceID.String())
	}
	return json.NewEncoder(w).Encode(res)
}

func EncodeListDataResponse(
	ctx context.Context,
	w http.ResponseWriter,
	data interface{},
) error {
	if f, ok := data.(endpoint_kit.Failer); ok {
		if e := f.Failed(); e != nil {
			EncodeError(ctx, e, w)
			return nil
		}
	}

	res := &endpoint.BasicResponse{
		Status: endpoint.StatusOk,
	}

	if d, ok := data.(endpoint.ExtraData); ok {
		res.Data = d.Data()
	} else {
		res.Data = data
	}

	if p, ok := data.(endpoint.Pager); ok {
		res.Paging = endpoint.NewPaging(p.Pagination())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	span := trace.FromContext(ctx)
	if span != nil {
		spanCtx := span.SpanContext()
		w.Header().Set("x-trace-id", spanCtx.TraceID.String())
	}
	return json.NewEncoder(w).Encode(res)
}

type File struct {
	Data []byte
	Name string
	Size int64
}

func EncodeFileResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	if f, ok := resp.(endpoint_kit.Failer); ok {
		if e := f.Failed(); e != nil {
			EncodeError(ctx, e, w)
			return nil
		}
	}

	file, ok := resp.(*File)
	if !ok {
		return errors.New("File not exist")
	}

	// Transmit the headers
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	w.Header().Set("Content-Length", strconv.FormatInt(file.Size, 10))
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write(file.Data)

	return nil
}

func EncodeErrorV2(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	res := &endpoint.BasicResponse{Status: endpoint.StatusError}
	statusCode := 400
	switch e := err.(type) {
	case endpoint.ErrBadRequest:
		res.Error = map[string]interface{}{
			"code":    "VALIDATION_ERROR",
			"message": e.Error(),
		}
	case endpoint.BizError:
		res.Error = map[string]interface{}{
			"code":    e.Code(),
			"message": e.Error(),
		}

		statusCode = http.StatusOK
	case endpoint.Error:
		res.Error = map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
		}
		if e.HttpStatusCode != 0 {
			statusCode = e.HttpStatusCode
		}
	default:
		fmt.Println("Error happen:", err)
		packet := raven.NewPacket(
			err.Error(),
			raven.NewException(err, raven.GetOrNewStacktrace(err, 2, 3, nil)),
			endpoint.HttpFromContext(ctx),
		)
		raven.Capture(packet, nil)
		res.Error = endpoint.ErrInternal
	}

	if data, ok := err.(endpoint.ExtraData); ok {
		res.Error["data"] = data.Data()
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func EncodeRawDataResponse(
	ctx context.Context,
	w http.ResponseWriter,
	data interface{},
) error {
	if f, ok := data.(endpoint_kit.Failer); ok {
		if e := f.Failed(); e != nil {
			EncodeError(ctx, e, w)
			return nil
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	span := trace.FromContext(ctx)
	if span != nil {
		spanCtx := span.SpanContext()
		w.Header().Set("x-trace-id", spanCtx.TraceID.String())
	}
	return json.NewEncoder(w).Encode(data)
}
