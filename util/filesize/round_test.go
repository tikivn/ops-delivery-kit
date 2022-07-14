package filesize

import "testing"

func TestHumanFileSize(t *testing.T) {
	type args struct {
		size float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1Gb",
			args: args{
				size: 1073741824,
			},
			want: "1 GB",
		},
		{
			name: "Test 1.5 Gb",
			args: args{
				size: 1610612736,
			},
			want: "1.5 GB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanFileSize(tt.args.size); got != tt.want {
				t.Errorf("HumanFileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
