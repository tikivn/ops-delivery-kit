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
		host:     "https://driver-tracking-uat.tiki.com.vn",
		httpDoer: &http.Client{
			Timeout: time.Minute,
		},
	}
	activeDrivers, err := c.GetActiveDriversInArea(context.Background(), 101, 101, 5000)
	assert.Nil(t, err)
	assert.NotEqualf(t, len(activeDrivers),0,"")
	fmt.Println(activeDrivers)
}
