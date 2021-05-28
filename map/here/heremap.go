package here

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// ErrNotFound is the error returns when we can't found lat/lng
var ErrNotFound = errors.New("Not found lat/lng of the given location")

type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client interface {
	Geocode(ctx context.Context, city string, district string, subdistrict string, street string) (*GeocodeResult, error)
}

type hereMapClient struct {
	host        string
	appID       string
	apiKey      string
	routingMode string
	httpDoer    HttpDoer
}

type envelope struct {
	Response Response `json:"Response"`
}

type Response struct {
	View []View `json:"View"`
}

type View struct {
	Result []Result `json:"Result"`
}

type Result struct {
	Location Location `json:"Location"`
}

type Location struct {
	DisplayPosition DisplayPosition `json:"DisplayPosition"`
}

type DisplayPosition struct {
	Lat float64 `json:"Latitude"`
	Lng float64 `json:"Longitude"`
}

func NewHereMapClient(heremapInfo map[string]string) Client {
	httpDoer := http.DefaultClient
	host := heremapInfo["host"]
	appID := heremapInfo["app_id"]
	apiKey := heremapInfo["apikey"]

	// Chỉ dùng cho routing (tính khoảng cách)
	routingMode := heremapInfo["routing_mode"]

	return &hereMapClient{
		host:        host,
		appID:       appID,
		apiKey:      apiKey,
		routingMode: routingMode,
		httpDoer:    httpDoer,
	}
}

type LatLng struct {
	Lat float64
	Lng float64
}

type GeocodeResult struct {
	PlusCode string
	LatLng   LatLng
}

func (h *hereMapClient) Geocode(ctx context.Context, city string, district string, subdistrict string, street string) (*GeocodeResult, error) {
	req, err := http.NewRequest("GET", h.host, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("app_id", h.appID)
	q.Add("apikey", h.apiKey)
	q.Add("country", "Việt Nam")
	q.Add("city", city)
	q.Add("district", district)
	q.Add("subdistrict", subdistrict)
	q.Add("street", street)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Content-Type", "application/json")
	req = req.WithContext(ctx)

	res, err := h.httpDoer.Do(req)

	if err != nil {
		return nil, errors.Wrapf(err, "Here Map: Response error for query")
	}
	defer res.Body.Close()

	e := &envelope{}
	if err = json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, errors.Wrapf(err, "Here Map: couldnt decode json")
	}

	if len(e.Response.View) > 0 {
		return &GeocodeResult{
			PlusCode: "",
			LatLng: LatLng{
				Lat: e.Response.View[0].Result[0].Location.DisplayPosition.Lat,
				Lng: e.Response.View[0].Result[0].Location.DisplayPosition.Lng,
			},
		}, nil
	} else {
		return nil, ErrNotFound
	}
}
