package endpoint

import (
	"testing"
)

func TestErrBadRequest_Error(t *testing.T) {
	type fields struct {
		Field   string
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "Full data",
			fields: fields{
				Field:   "a",
				Message: "b",
			},
			want: "a: b",
		},
		{
			name: "Missing field",
			fields: fields{
				Message: "m",
			},
			want: "m",
		},
		{
			name: "Missing message",
			fields: fields{
				Field: "a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrBadRequest{
				Field:   tt.fields.Field,
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
