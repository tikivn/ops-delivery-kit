package endpoint

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
