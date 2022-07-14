package algorithm

import (
	"reflect"
	"testing"
)

func Test_RemoveIndexString(t *testing.T) {
	type args struct {
		slice []string
		index int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "happy case 01",
			args: args{
				slice: []string{"a", "b", "c"},
				index: 1,
			},
			want: []string{"a", "c"},
		},
		{
			name: "happy case 02",
			args: args{
				slice: []string{"a", "b", "c"},
				index: 2,
			},
			want: []string{"a", "b"},
		},
		{
			name: "happy case 03",
			args: args{
				slice: []string{"a", "b", "c"},
				index: 3,
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "happy case 04",
			args: args{
				slice: []string{"a", "b", "c"},
				index: 0,
			},
			want: []string{"b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveIndexString(tt.args.slice, tt.args.index)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ReverseStringSlice(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty slice",
			args: args{s: []string{}},
			want: []string{},
		},
		{
			name: "slice has 1 element",
			args: args{s: []string{"a"}},
			want: []string{"a"},
		},
		{
			name: "slice has 2 elements",
			args: args{s: []string{"a", "b"}},
			want: []string{"b", "a"},
		},
		{
			name: "slice has more than 2 elements",
			args: args{s: []string{"a", "b", "d", "z"}},
			want: []string{"z", "d", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseStringSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
