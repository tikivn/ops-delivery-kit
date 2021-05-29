package luffy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"go.opencensus.io/plugin/ochttp"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client interface {
	CreateShipment(ctx context.Context, req Request) (CreateShipmentResult, error)
	GetQuotes(ctx context.Context, req Request) (QuotesResult, error)
}

type client struct {
	host     string
	httpDoer HttpDoer
}

var _ Client = (*client)(nil)

func NewLuffyClient(host string, httpDoer HttpDoer) *client {
	if httpDoer == nil {
		httpDoer = &http.Client{
			Transport: &ochttp.Transport{},
			Timeout:   10 * time.Second,
		}
	}

	return &client{
		host:     host,
		httpDoer: httpDoer,
	}
}

type shipcodeResponse struct {
	Status   string `json:"status"`
	Shipcode string `json:"shipcode"`
	Error    string `json:"error"`
}

func (c *client) GetQuotes(ctx context.Context, payload Request) (QuotesResult, error) {
	path := fmt.Sprintf("%v/v1/request/quotes", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "luffy: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "luffy: Cant get quotes")
	}
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "luffy: Response error for query")
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := QuotesResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return e, errors.Wrapf(err, "luffy: couldnt decode json, body %s", string(body))
		}

		return e, nil
	}

	return QuotesResult{}, errors.Errorf("luffy: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) CreateShipment(ctx context.Context, payload Request) (CreateShipmentResult, error) {
	path := fmt.Sprintf("%v/v1/request/create", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "luffy: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "luffy: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "luffy: Response error for query")
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := CreateShipmentResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return e, errors.Wrapf(err, "luffy: couldnt decode json, body %s", string(body))
		}

		return e, nil
	}

	return CreateShipmentResult{}, errors.Errorf("luffy: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}
