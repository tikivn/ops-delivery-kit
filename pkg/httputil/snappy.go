package httputil

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golang/snappy"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type closer struct {
	io.Closer
	io.Reader
}

type snappySupport struct {
	HttpDoer
}

func NewSnappySupport(httpDoer HttpDoer) HttpDoer {
	return &snappySupport{
		HttpDoer: httpDoer,
	}
}

func (s *snappySupport) Do(req *http.Request) (*http.Response, error) {
	res, err := s.HttpDoer.Do(req)
	if err == nil {
		isSnappy := res.Header.Get("compression") == "snappy" ||
			res.Header.Get("Content-Encoding") == "x-snappy-framed"
		isFramed := res.Header.Get("Content-Encoding") == "x-snappy-framed"
		if isSnappy {
			if isFramed {
				res.Body = &closer{
					Closer: res.Body,
					Reader: snappy.NewReader(res.Body),
				}
			} else {
				defer res.Body.Close()
				src, err := ioutil.ReadAll(res.Body)
				if err != nil {
					return res, err
				}

				dst, err := snappy.Decode(nil, src)
				if err != nil {
					return res, err
				}

				res.Body = ioutil.NopCloser(bytes.NewReader(dst))
			}
		}
	}

	return res, err
}
