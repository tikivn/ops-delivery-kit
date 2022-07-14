package themis

import (
	"context"
	"testing"
)

func Test_client_AuthorizeToken(t *testing.T) {
	type fields struct {
		host       string
		httpClient HttpDoer
		token      TokenProvider
	}
	type args struct {
		ctx      context.Context
		resource string
		action   string
		token    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				host:       tt.fields.host,
				httpClient: tt.fields.httpClient,
				token:      tt.fields.token,
			}
			got, err := c.AuthorizeToken(tt.args.ctx, tt.args.resource, tt.args.action, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AuthorizeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.AuthorizeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
