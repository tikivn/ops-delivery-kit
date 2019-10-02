package oms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	defaultHost = "http://oms.tiki.services"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client interface {
	FetchOrder(ctx context.Context, code string) (*Order, error)
}

type client struct {
	host     string
	httpDoer HttpDoer
}

func NewClient(httpDoer HttpDoer) (Client, error) {
	return NewClientWithHost(defaultHost, httpDoer)
}

func NewClientWithHost(host string, httpDoer HttpDoer) (Client, error) {
	if httpDoer == nil {
		httpDoer = http.DefaultClient
	}

	return &client{
		host:     host,
		httpDoer: httpDoer,
	}, nil
}

func (c *client) FetchOrder(ctx context.Context, code string) (*Order, error) {
	path := fmt.Sprintf("%s/v3/orders/%s", c.host, code)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "OMS: (%s) Cant create Request", code)
	}
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "OMS: (%s) response error for order", code)
	}
	defer res.Body.Close()

	e := &envelope{}
	if err = json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, errors.Wrapf(err, "OMS: (%s) couldnt decode json", code)
	}
	if e.Err != nil {
		return nil, errors.Wrap(e.Err, "OMS")
	}
	return e.Order, nil
}

type envelope struct {
	Order *Order    `json:"order"`
	Err   *omsError `json:"error"`
}
type omsError struct {
	Message string `json:"message"`
}

func (e omsError) Error() string {
	return e.Message
}
