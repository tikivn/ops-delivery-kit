package luffy

import (
	"context"
	"reflect"
	"testing"
)

func Test_client_GetQuotes(t *testing.T) {
	type fields struct {
		host     string
		httpDoer HttpDoer
	}
	type args struct {
		ctx     context.Context
		payload Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []QuotesResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				host:     tt.fields.host,
				httpDoer: tt.fields.httpDoer,
			}
			got, err := c.GetQuotes(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetQuotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.GetQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
