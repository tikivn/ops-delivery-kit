package util

import (
	"testing"
)

func TestUcFirst(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{
				str: "normal",
			},
			want: "Normal",
		},
		{
			name: "underscore",
			args: args{
				str: "under_score",
			},
			want: "Under_score",
		},
		{
			name: "dash",
			args: args{
				str: "dash-dash",
			},
			want: "Dash-dash",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UcFirst(tt.args.str); got != tt.want {
				t.Errorf("UcFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveUnicodeNull(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with null",
			args: args{
				str: "With null\u0000",
			},
			want: "With null",
		},
		{
			name: "with 2 null",
			args: args{
				str: "With null \u0000 \u0000",
			},
			want: "With null  ",
		},
		{
			name: "with null in word",
			args: args{
				str: "Wi\u0000th null \x00",
			},
			want: "With null ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveUnicodeNull(tt.args.str); got != tt.want {
				t.Errorf("RemoveUnicodeNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToFloat64(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "with string",
			args: args{
				data: "8.3",
			},
			want: 8.3,
		},
		{
			name: "with float64",
			args: args{
				data: 1000.00,
			},
			want: 1000.00,
		},
		{
			name: "with text",
			args: args{
				data: "text",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "with other",
			args: args{
				data: []string{"slice"},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "with nil",
			args: args{
				data: nil,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "with empty string",
			args: args{
				data: "",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToFloat64(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InterfaceToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseValidForHereMap(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				addr: "Phường 05",
			},
			want: "phường 5",
		},
		{
			name: "test 2",
			args: args{
				addr: "Phường 11",
			},
			want: "phường 11",
		},
		{
			name: "test 3",
			args: args{
				addr: "Phường 5",
			},
			want: "phường 5",
		},
		{
			name: "test 4",
			args: args{
				addr: "Phường Đập Đá",
			},
			want: "Phường Đập Đá",
		},
		{
			name: "test 5",
			args: args{
				addr: "   Phường 05   ",
			},
			want: "phường 5",
		},
		{
			name: "test 6",
			args: args{
				addr: "Quận 04",
			},
			want: "quận 4",
		},
		{
			name: "test 7",
			args: args{
				addr: "Quận 10",
			},
			want: "quận 10",
		},
		{
			name: "test 8",
			args: args{
				addr: "Quận 5",
			},
			want: "quận 5",
		},
		{
			name: "test 9",
			args: args{
				addr: "Quận Bình Thạnh",
			},
			want: "Quận Bình Thạnh",
		},
		{
			name: "test 10",
			args: args{
				addr: "   Quận 04   ",
			},
			want: "quận 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseValidForHereMap(tt.args.addr); got != tt.want {
				t.Errorf("ParseValidSubdistrictForHereMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractTikiCode(t *testing.T) {
	type args struct {
		wardTikiCode string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			args: args{
				wardTikiCode: "VN039007012",
			},
			want:  "VN039007",
			want1: "VN039",
		},
		{
			name: "wrong 1",
			args: args{
				wardTikiCode: "VN03900",
			},
			want:  "",
			want1: "VN039",
		},
		{
			name: "wrong 2",
			args: args{
				wardTikiCode: "VN03",
			},
			want:  "",
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ExtractTikiCode(tt.args.wardTikiCode)
			if got != tt.want {
				t.Errorf("ExtractTikiCode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtractTikiCode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
