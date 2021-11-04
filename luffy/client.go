package luffy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client interface {
	GetQuotes(ctx context.Context, req Request) (QuotesResult, error)
	GenerateShipcode(ctx context.Context, payload Payload) (string, error)
	CreateShipment(ctx context.Context, req Request) (CreateShipmentResult, error)
	CancelShipment(ctx context.Context, trackingInfo TrackingInfo) (bool, error)
	GetShipment(ctx context.Context, trackingInfo TrackingInfo) (InfoResult, error)
}

type client struct {
	host     string
	httpDoer HttpDoer
	clientID string
}

var _ Client = (*client)(nil)

func NewLuffyClient(host string, clientID string, httpDoer HttpDoer) *client {
	if httpDoer == nil {
		httpDoer = &http.Client{
			Transport: &ochttp.Transport{},
			Timeout:   10 * time.Second,
		}
	}

	return &client{
		host:     host,
		httpDoer: httpDoer,
		clientID: clientID,
	}
}

type shipcodeResponse struct {
	Status   string `json:"status"`
	Shipcode string `json:"shipcode"`
	Error    string `json:"error"`
}

func (c *client) GenerateShipcode(ctx context.Context, payload Payload) (string, error) {
	path := fmt.Sprintf("%v/v1/shipcode/generate", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrapf(err, "luffy: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return "", errors.Wrapf(err, "luffy: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "luffy: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := shipcodeResponse{}

		if err := json.Unmarshal(result, &e); err != nil {
			return "", errors.Wrapf(err, "luffy: couldnt decode json, body %s, response %s", string(body), result)
		}

		return e.Shipcode, nil
	}

	return "", errors.Errorf("luffy: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
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
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "luffy: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := CreateShipmentResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return e, errors.Wrapf(err, "luffy: couldnt decode json, body %s, response %s", string(body), result)
		}

		return e, nil
	}

	failedResponse := FailedResponse{}
	if err = json.Unmarshal(result, &failedResponse); err == nil {
		return CreateShipmentResult{}, CreateShipmentFailError{Message: failedResponse.Error}
	}

	return CreateShipmentResult{}, errors.Errorf("luffy: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) CancelShipment(ctx context.Context, trackingInfo TrackingInfo) (bool, error) {
	path := fmt.Sprintf("%v/v1/request/cancel", c.host)

	body, err := json.Marshal(trackingInfo)
	if err != nil {
		return false, errors.Wrapf(err, "luffy: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return false, errors.Wrapf(err, "luffy: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return false, errors.Wrapf(err, "luffy: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := CreateShipmentResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return false, errors.Wrapf(err, "luffy: couldnt decode json, body %s, response %s", string(body), result)
		}

		return true, nil
	}

	failedResponse := FailedResponse{}
	if err = json.Unmarshal(result, &failedResponse); err == nil {
		return false, CancelShipmentFailError{Message: failedResponse.Error}
	}

	return false, errors.Errorf("luffy: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) GetShipment(ctx context.Context, trackingInfo TrackingInfo) (InfoResult, error) {
	path := fmt.Sprintf("%v/v1/request/getInfo", c.host)

	body, err := json.Marshal(trackingInfo)
	if err != nil {
		return InfoResult{}, errors.Wrapf(err, "luffy: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return InfoResult{}, errors.Wrapf(err, "luffy: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return InfoResult{}, errors.Wrapf(err, "luffy: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := InfoResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return InfoResult{}, errors.Wrapf(err, "luffy: couldnt decode json, body %s, response %s", string(body), result)
		}

		return e, nil
	}

	return InfoResult{}, errors.Errorf("luffy: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) GetQuotes(ctx context.Context, payload Request) (QuotesResult, error) {
	path := fmt.Sprintf("%v/v2/request/quotes", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "luffy: Can not marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "luffy: Can not get quotes")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "luffy: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := QuotesResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return e, errors.Wrapf(err, "luffy: couldnt decode json, body %s, response %s", string(body), result)
		}

		return e, nil
	}

	failedResponse := FailedResponse{}
	if err = json.Unmarshal(result, &failedResponse); err == nil {
		return QuotesResult{}, GetQuotesFailError{Message: failedResponse.Error}
	}

	return QuotesResult{}, errors.Errorf("luffy: server response status code = %d, response = %s", res.StatusCode, result)
}
