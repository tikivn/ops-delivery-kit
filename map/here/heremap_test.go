package here

import (
	"context"
	"reflect"
	"testing"
)

func Test_hereMapClient_Geocode(t *testing.T) {
	// defaultHttpDoer := &http.Client{
	// 	Timeout: time.Second * 15,
	// }

	type fields struct {
		host        string
		appID       string
		apiKey      string
		routingMode string
		httpDoer    HttpDoer
	}
	type args struct {
		ctx         context.Context
		city        string
		district    string
		subdistrict string
		street      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GeocodeResult
		wantErr bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hereMapClient{
				host:        tt.fields.host,
				appID:       tt.fields.appID,
				apiKey:      tt.fields.apiKey,
				routingMode: tt.fields.routingMode,
				httpDoer:    tt.fields.httpDoer,
			}
			got, err := h.Geocode(tt.args.ctx, tt.args.city, tt.args.district, tt.args.subdistrict, tt.args.street)
			if (err != nil) != tt.wantErr {
				t.Errorf("hereMapClient.Geocode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hereMapClient.Geocode() = %v, want %v", got, tt.want)
			}
		})
	}
}
