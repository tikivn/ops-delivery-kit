package endpoint

import (
	"strings"
)

type BizError interface {
	error
	Code() string
}

type Error struct {
	Code           string `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"-"`
}

func (e Error) Error() string {
	return e.Message
}

var ErrInternal = map[string]interface{}{
	"code":    "INTERNAL_ERROR",
	"message": "Có lỗi hệ thống xảy ra",
}

type ErrBadRequest struct {
	Field   string
	Message string
}

func (e ErrBadRequest) Error() string {
	arr := make([]string, 0, 2)
	if str := strings.TrimSpace(e.Field); str != "" {
		arr = append(arr, str)
	}

	if str := strings.TrimSpace(e.Message); str != "" {
		arr = append(arr, str)
	}

	return strings.Join(arr, ": ")
}

func (e ErrBadRequest) Failed() error {
	return e
}
