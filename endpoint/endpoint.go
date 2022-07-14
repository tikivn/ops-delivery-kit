package endpoint

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/getsentry/raven-go"
	"go.opencensus.io/trace"

	"github.com/go-kit/kit/endpoint"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 20

	PageSizeQuery = "size"
	PageNumQuery  = "page"
)

type ErrBadRequest struct {
	Field   string
	Message string
}

func (e ErrBadRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type Status string

const (
	StatusOk    = Status("OK")
	StatusError = Status("ERROR")
)

type pagingResponse struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	PerPage     int `json:"per_page"`
	LastPage    int `json:"last_page"`
}

func (p *pagingResponse) Pagination() (int, int, int) {
	return p.Total, p.CurrentPage, p.PerPage
}

func NewPaging(total, page, size int) *pagingResponse {
	if total == 0 {
		return &pagingResponse{
			PerPage: size,
		}
	}

	return &pagingResponse{
		Total:       total,
		CurrentPage: page,
		From:        (page-1)*size + 1,
		To:          page * size,
		PerPage:     size,
		LastPage: func(total, perpage int) int {
			if total%perpage != 0 {
				return total/perpage + 1
			}
			return total / perpage
		}(total, size),
	}
}

type BasicResponse struct {
	Status Status                 `json:"status"`
	Data   interface{}            `json:"data,omitempty"`
	Paging *pagingResponse        `json:"paging,omitempty"`
	Error  map[string]interface{} `json:"error,omitempty"`
}

type ExtraData interface {
	Data() interface{}
}

var errInternal = map[string]interface{}{
	"code":    "INTERNAL_ERROR",
	"message": "Có lỗi hệ thống xảy ra",
}

// EncodeError encode errors from business-logic
func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	res := &BasicResponse{Status: StatusError}
	var statusCode int
	switch e := err.(type) {
	case ErrBadRequest:
		res.Error = map[string]interface{}{
			"code":    "VALIDATION_ERROR",
			"message": e.Error(),
		}

		statusCode = http.StatusBadRequest
	case BizError:
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
			HttpFromContext(ctx),
		)
		raven.Capture(packet, nil)
		res.Error = errInternal

		statusCode = http.StatusInternalServerError
	}

	if data, ok := err.(ExtraData); ok {
		res.Error["data"] = data.Data()
	}

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(res)
}

type Pager interface {
	Pagination() (total, page, size int)
}

func EncodeResponse(
	ctx context.Context,
	w http.ResponseWriter,
	data interface{},
) error {
	if f, ok := data.(endpoint.Failer); ok {
		if e := f.Failed(); e != nil {
			EncodeError(ctx, e, w)
			return nil
		}
	}

	res := &BasicResponse{
		Status: StatusOk,
		Data:   data,
	}

	if p, ok := data.(Pager); ok {
		res.Paging = NewPaging(p.Pagination())
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
	if f, ok := data.(endpoint.Failer); ok {
		if e := f.Failed(); e != nil {
			EncodeError(ctx, e, w)
			return nil
		}
	}

	res := &BasicResponse{
		Status: StatusOk,
	}

	if d, ok := data.(ExtraData); ok {
		res.Data = d.Data()
	} else {
		res.Data = data
	}

	if p, ok := data.(Pager); ok {
		res.Paging = NewPaging(p.Pagination())
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
	if f, ok := resp.(endpoint.Failer); ok {
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
	res := &BasicResponse{Status: StatusError}
	statusCode := 400
	switch e := err.(type) {
	case ErrBadRequest:
		res.Error = map[string]interface{}{
			"code":    "VALIDATION_ERROR",
			"message": e.Error(),
		}
	case BizError:
		res.Error = map[string]interface{}{
			"code":    e.Code(),
			"message": e.Error(),
		}

		statusCode = http.StatusOK
	case Error:
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
			HttpFromContext(ctx),
		)
		raven.Capture(packet, nil)
		res.Error = errInternal
	}

	if data, ok := err.(ExtraData); ok {
		res.Error["data"] = data.Data()
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
}

func EncodeRawDataResponse(
	ctx context.Context,
	w http.ResponseWriter,
	data interface{},
) error {
	if f, ok := data.(endpoint.Failer); ok {
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

type PagingQuerier struct {
	PageNum  int
	PageSize int
}

func (p PagingQuerier) IsZero() bool {
	return p.PageSize == 0 && p.PageNum == 0
}

func PagingFromRequest(_ context.Context, r *http.Request) (PagingQuerier, error) {
	q := r.URL.Query()

	pageNum := 0
	pageNumStr := strings.TrimSpace(q.Get(PageNumQuery))
	if pageNumStr != "" {
		n, err := strconv.ParseInt(pageNumStr, 10, 64)
		if err != nil {
			return PagingQuerier{}, err
		}
		pageNum = int(n)
	}

	pageSize := 0
	pageSizeStr := strings.TrimSpace(q.Get(PageSizeQuery))
	if pageSizeStr != "" {
		n, err := strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil {
			return PagingQuerier{}, err
		}
		pageSize = int(n)
	}

	return PagingQuerier{
		PageNum:  pageNum,
		PageSize: pageSize,
	}, nil
}

func PagingWithDefault(ctx context.Context, r *http.Request) (PagingQuerier, error) {
	pager, err := PagingFromRequest(ctx, r)
	if err != nil {
		return PagingQuerier{}, err
	}

	if !pager.IsZero() {
		return pager, nil
	}

	return PagingQuerier{
		PageNum:  DefaultPage,
		PageSize: DefaultPageSize,
	}, nil
}
