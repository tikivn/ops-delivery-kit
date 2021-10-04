package tms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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

	GetActiveP2PDrivers(ctx context.Context, tikiCode string, teamCodes []string) (driverIds []string, err error)
}

type client struct {
	host     string
	httpDoer HttpDoer
	clientID string
}

var _ Client = (*client)(nil)

func NewTMSClient(host string, clientID string, httpDoer HttpDoer) *client {
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
	path := fmt.Sprintf("%v/s2s/express/generateTrackingCode", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrapf(err, "tms: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return "", errors.Wrapf(err, "tms: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "tms: Response error for query")
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
			return "", errors.Wrapf(err, "tms: couldnt decode json, body %s, response %s", string(body), result)
		}

		if e.Error != "" {
			return "", errors.New(e.Error)
		}

		return e.Shipcode, nil
	}

	return "", errors.Errorf("tms: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) CreateShipment(ctx context.Context, payload Request) (CreateShipmentResult, error) {
	path := fmt.Sprintf("%v/s2s/express/createShipment", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "tms: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "tms: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return CreateShipmentResult{}, errors.Wrapf(err, "tms: Response error for query")
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
			return e, errors.Wrapf(err, "tms: couldnt decode json, body %s, response %s", string(body), result)
		}

		if e.Error != "" {
			return CreateShipmentResult{}, CreateShipmentFailError{Message: e.Error}
		}

		return e, nil
	}

	return CreateShipmentResult{}, errors.Errorf("tms: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) CancelShipment(ctx context.Context, trackingInfo TrackingInfo) (bool, error) {
	path := fmt.Sprintf("%v/s2s/express/cancelShipment", c.host)

	body, err := json.Marshal(trackingInfo)
	if err != nil {
		return false, errors.Wrapf(err, "tms: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return false, errors.Wrapf(err, "tms: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return false, errors.Wrapf(err, "tms: Response error for query")
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
			return false, errors.Wrapf(err, "tms: couldnt decode json, body %s, response %s", string(body), result)
		}

		if e.Error != "" {
			return false, errors.New(e.Error)
		}

		return true, nil
	}

	return false, errors.Errorf("tms: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) GetShipment(ctx context.Context, trackingInfo TrackingInfo) (InfoResult, error) {
	path := fmt.Sprintf("%v/s2s/express/tracking", c.host)

	body, err := json.Marshal(trackingInfo)
	if err != nil {
		return InfoResult{}, errors.Wrapf(err, "tms: Cant marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return InfoResult{}, errors.Wrapf(err, "tms: Cant create Request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return InfoResult{}, errors.Wrapf(err, "tms: Response error for query")
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
			return InfoResult{}, errors.Wrapf(err, "tms: couldnt decode json, body %s, response %s", string(body), result)
		}

		if e.Error != "" {
			return InfoResult{}, errors.New(e.Error)
		}

		return e, nil
	}

	return InfoResult{}, errors.Errorf("tms: server response status code = %d, payload = %+v, response = %s", res.StatusCode, string(body), result)
}

func (c *client) GetQuotes(ctx context.Context, payload Request) (QuotesResult, error) {
	path := fmt.Sprintf("%v/s2s/express/quotes", c.host)

	body, err := json.Marshal(payload)
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "tms: Can not marshal request body")
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "tms: Can not get quotes")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return QuotesResult{}, errors.Wrapf(err, "tms: Response error for query")
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
			return e, errors.Wrapf(err, "tms: couldnt decode json, body %s, response %s", string(body), result)
		}

		if e.Error != "" {
			return QuotesResult{}, GetQuotesFailError{Message: e.Error}
		}

		return e, nil
	}

	return QuotesResult{}, errors.Errorf("tms: server response status code = %d, response = %s", res.StatusCode, result)
}

func (c *client) GetActiveP2PDrivers(ctx context.Context, tikiCode string, teamCodes []string) (driverIds []string, err error) {
	path := fmt.Sprintf("%v/s2s/v1/attendance/active-p2p-drivers?tiki_code=%s", c.host, tikiCode)
	if len(teamCodes) > 0 {
		path += fmt.Sprintf("&team_codes=%s", strings.Join(teamCodes, ","))
	}

	body := []byte(nil)

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))
	if err != nil {
		return []string{}, errors.Wrapf(err, "tms: Can not get active drivers")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return []string{}, errors.Wrapf(err, "tms: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		e := ActiveDriversResult{}

		if err := json.Unmarshal(result, &e); err != nil {
			return []string{}, errors.Wrapf(err, "tms: couldnt decode json, body %s, response %s", string(body), result)
		}

		if e.Status == "ERROR" {
			return []string{}, GetQuotesFailError{Message: e.Error.Message}
		}

		return e.Data.DriverIds, nil
	}

	return []string{}, errors.Errorf("tms: server response status code = %d, response = %s", res.StatusCode, result)
}
