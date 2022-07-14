package endpoint

import "fmt"

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
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
