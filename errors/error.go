package errors

import "github.com/pkg/errors"

type errBusiness struct {
	code    string
	message string
	data    interface{}
}

func CreateBusinessError(code, message string, data interface{}) *errBusiness {
	return &errBusiness{
		code:    code,
		message: message,
		data:    data,
	}
}

func (e *errBusiness) Code() string {
	return e.code
}

func (e *errBusiness) Error() string {
	return e.message
}

func (e *errBusiness) Data() interface{} {
	return e.data
}

func (e *errBusiness) Failed() error {
	return errors.New(e.message)
}
