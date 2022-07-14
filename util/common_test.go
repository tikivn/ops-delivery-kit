package util

import (
	"testing"
)

func Test_InArrayString(t *testing.T) {
	type args struct {
		arr  []string
		item string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Valid case
		{
			name: "Happy case 01",
			args: args{
				arr:  []string{"a", "b"},
				item: "a",
			},
			want: true,
		},
		{
			name: "Happy case 02",
			args: args{
				arr:  []string{"a", "b"},
				item: "b",
			},
			want: true,
		},
		// Invalid case
		{
			name: "Unhappy case 01",
			args: args{
				arr:  []string{"a", "b"},
				item: "",
			},
			want: false,
		},
		{
			name: "Unhappy case 01",
			args: args{
				arr:  []string{"a", "b"},
				item: "c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InArrayString(tt.args.arr, tt.args.item)
			if got != tt.want {
				t.Errorf("InArrayString() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
