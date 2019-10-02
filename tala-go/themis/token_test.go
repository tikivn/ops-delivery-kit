package themis

import (
	"context"
	"net/http"
	"testing"
)

func TestClientToken_AccessToken(t *testing.T) {
	type fields struct {
		host         string
		clientID     string
		clientSecret string
		http         HttpDoer
		token        *AccessToken
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "x",
			fields: fields{
				host:         "https://api.tala.xyz",
				clientID:     "trn:tiki:tms",
				clientSecret: "m.s0wllvRDCE",
				http:         http.DefaultClient,
			},
			args: args{
				in0: context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ClientToken{
				host:         tt.fields.host,
				clientID:     tt.fields.clientID,
				clientSecret: tt.fields.clientSecret,
				http:         tt.fields.http,
				token:        tt.fields.token,
			}
			got, err := c.AccessToken(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientToken.AccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientToken.AccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
