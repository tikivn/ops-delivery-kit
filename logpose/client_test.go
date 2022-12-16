package logpose

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// go test -count=1 ./logpose -run Test_client_GetActiveDriversInArea -v
func Test_client_GetActiveDriversInArea(t *testing.T) {

	c := &client{
		host: "https://driver-tracking-uat.tiki.com.vn",
		httpDoer: &http.Client{
			Timeout: time.Minute,
		},
	}
	activeDrivers, err := c.GetActiveDriversInArea(context.Background(), 101, 101, 5000)
	assert.Nil(t, err)
	assert.NotEqualf(t, len(activeDrivers), 0, "")
	fmt.Println(activeDrivers)
}

// go test -count=1 ./logpose -run Test_client_Geocode -v
func Test_client_Geocode(t *testing.T) {
	c := &client{
		host: "https://driver-tracking-uat.tiki.com.vn",
		httpDoer: &http.Client{
			Timeout: time.Minute,
		},
	}
	activeDrivers, err := c.Geocode(context.Background(), GeocodePayload{
		Province: "Hà Nội",
		District: "Quận Hoàng Mai",
		Ward:     "Phường Mai Động",
		Street:   "89 Lĩnh Nam",
	})
	assert.Nil(t, err)
	assert.Equalf(t, Coordinates{
		Lat: 20.9884168,
		Lng: 105.8665135,
	}, activeDrivers, "")
	fmt.Println(activeDrivers)
}
