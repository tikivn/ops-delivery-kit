package endpoint

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 20

	PageSizeQuery = "size"
	PageNumQuery  = "page"
)

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

type Pager interface {
	Pagination() (total, page, size int)
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

type Pagable interface {
	ExtraData
	Pager
}

type pagableResponse struct {
	ls                interface{}
	total, page, size int
}

func (r pagableResponse) Data() interface{} {
	return r.ls
}

func (r pagableResponse) Pagination() (total, page, size int) {
	return r.total, r.page, r.size
}

func NewPagingResponse(any interface{}, total, page, size int) Pagable {
	if page <= 0 && size <= 0 && total > 0 {
		page = 1
		size = total
	}
	return pagableResponse{
		ls:    any,
		total: total,
		page:  page,
		size:  size,
	}
}
