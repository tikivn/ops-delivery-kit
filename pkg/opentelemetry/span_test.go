package opentelemetry

import (
	"reflect"
	"testing"
)

func Test_newSpanFuncConfig(t *testing.T) {
	type args struct {
		options []SpanFunctionOption
	}
	tests := []struct {
		name string
		args args
		want spanFunctionConfig
	}{
		// TODO: Add test cases.
		{
			name: "Test happy case",
			args: args{
				options: []SpanFunctionOption{WithCaller("caller")},
			},
			want: spanFunctionConfig{
				caller: "caller",
			},
		},
		{
			name: "Test happy case with multiple option",
			args: args{
				options: []SpanFunctionOption{
					WithCaller("caller"),
					WithIdentifier("identify"),
				},
			},
			want: spanFunctionConfig{
				caller:   "caller",
				identify: "identify",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSpanFuncConfig(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSpanFuncConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
