package logpose

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client interface {
	GetActiveDriversInArea(ctx context.Context, lat, lng float32, distance int) ([]ActiveDriversResult, error)
}

type client struct {
	host     string
	httpDoer HttpDoer
	clientID string
}

var _ Client = (*client)(nil)

func NewLogPoseClient(host string, clientID string, httpDoer HttpDoer) *client {
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

func (c *client) GetActiveDriversInArea(ctx context.Context, lat, lng float32, distance int) ([]ActiveDriversResult, error) {
	var e []ActiveDriversResult

	path := fmt.Sprintf("%s/v1/driver-actives-in-area?lat=%f&lng=%f&distance=%d", c.host, lat, lng, distance)

	body := []byte(nil)

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(body))
	if err != nil {
		return e, errors.Wrapf(err, "tms: Can not get active drivers")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", c.clientID)
	req = req.WithContext(ctx)

	res, err := c.httpDoer.Do(req)
	if err != nil {
		return e, errors.Wrapf(err, "logpose: Response error for query")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(res.Body)

	result, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {

		if err := json.Unmarshal(result, &e); err != nil {
			return e, errors.Wrapf(err, "logpose: couldnt decode json, body %s, response %s", string(body), result)
		}

		return e, nil
	}

	return e, errors.Errorf("logpose: server response status code = %d, response = %s", res.StatusCode, result)
}
