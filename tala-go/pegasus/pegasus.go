package pegasus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

const (
	HostQueryProduction  = "http://pegasus-query.tiki.services"
	HostSingleProduction = "http://pegasus.tiki.services"
	HostQueryStaging     = "http://pegasus.dev.tiki.services"
	HostSingleStaging    = "http://pegasus.dev.tiki.services"
)

type Client interface {
	SingleProduct(ctx context.Context, productID int64) (*Product, error)
	QueryProducts(ctx context.Context, productIDs []int64) ([]Product, error)
}

type client struct {
	hostQuery  string
	hostSingle string
	httpDoer   HttpDoer
}

func NewClient(hostQuery, hostSingle string, httpDoer HttpDoer) Client {
	if httpDoer == nil {
		httpDoer = http.DefaultClient
	}

	return &client{
		hostQuery:  hostQuery,
		hostSingle: hostSingle,
		httpDoer:   NewSnappySupport(httpDoer),
	}
}

func (c *client) SingleProduct(ctx context.Context, productID int64) (*Product, error) {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "%s/v1/products/%d", c.hostSingle, productID)
	req, err := http.NewRequest("GET", builder.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Pegasus: new request error")
	}

	req = req.WithContext(ctx)
	res, err := c.httpDoer.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Pegasus: Query product error")
	}
	defer res.Body.Close()

	var e envelope
	err = json.NewDecoder(res.Body).
		Decode(&e)
	if err != nil {
		return nil, errors.Wrap(err, "Pegasus: decode json error")
	}
	if e.Err != nil {
		return nil, errors.Wrap(e.Err, "Pegasus")
	}
	return &e.Product, nil
}

func (c *client) QueryProducts(ctx context.Context, productIDs []int64) ([]Product, error) {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "%s/v1/products?ids=", c.hostQuery)
	for idx, pid := range productIDs {
		if idx != 0 {
			builder.WriteRune(',')
		}
		fmt.Fprintf(&builder, "%d", pid)
	}
	req, err := http.NewRequest("GET", builder.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Pegasus: new request error")
	}

	req = req.WithContext(ctx)
	res, err := c.httpDoer.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Pegasus: Query product error")
	}
	defer res.Body.Close()

	var products []Product
	err = json.NewDecoder(res.Body).
		Decode(&products)
	if err != nil {
		return nil, errors.Wrap(err, "Pegasus: decode json error")
	}
	return products, nil
}

type envelope struct {
	Product
	Err *pegasusError `json:"error"`
}
type pegasusError struct {
	Message string `json:"message"`
}

func (e pegasusError) Error() string {
	return e.Message
}
