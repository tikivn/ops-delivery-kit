package util

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCodeToUUID(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want uuid.UUID
	}{
		// TODO: Add test cases.
		{
			name: "Test exchagne key",
			args: args{
				code: "417399864-C",
			},
			want: uuid.MustParse("62ae8513-8762-4d89-ae8c-9b5c448bda32"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CodeToUUID(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CodeToUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
