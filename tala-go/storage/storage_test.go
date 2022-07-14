package storage_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/tikivn/ops-delivery-kit/tala-go/storage"
)

func Test_service_UploadFile(t *testing.T) {
	type args struct {
		ctx      context.Context
		bucket   string
		filename string
		file     io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *storage.UploadedFile
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:      context.Background(),
				bucket:   "tmp",
				filename: "happy_case.txt",
				file:     bytes.NewBufferString("happy case"),
			},
		},
	}
	s := storage.NewFileService(
		"http://uat.storage.tiki.services",
		http.DefaultTransport,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.UploadFile(tt.args.ctx, tt.args.bucket, tt.args.filename, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UploadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
