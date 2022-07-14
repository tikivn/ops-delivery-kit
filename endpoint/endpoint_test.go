package endpoint

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestPagingFromRequest(t *testing.T) {
	type args struct {
		in0 context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    PagingQuerier
		wantErr bool
	}{
		{
			name: "Test happy case",
			args: args{
				in0: nil,
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "page=1&size=10",
					},
				},
			},
			want: PagingQuerier{
				PageNum:  1,
				PageSize: 10,
			},
			wantErr: false,
		},
		{
			name: "Test case parse error",
			args: args{
				in0: nil,
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "page=1&size=10fds",
					},
				},
			},
			want:    PagingQuerier{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PagingFromRequest(tt.args.in0, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("PagingFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PagingFromRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
