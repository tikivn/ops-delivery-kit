package util

import "testing"

func TestFuzzyDecision(t *testing.T) {
	type args struct {
		code         string
		clientPrefix []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
		{
			name: "Test MBP case",
			args: args{
				code: "MBP/2021/10/045580",
				clientPrefix: []string{
					"MBP",
				},
			},
			want:  "MBP/2021/10/045580",
			want1: "",
		},
		{
			name: "Test normal order case",
			args: args{
				code: "123456",
				clientPrefix: []string{
					"MBP",
				},
			},
			want:  "123456",
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FuzzyDecision(tt.args.code, tt.args.clientPrefix)
			if got != tt.want {
				t.Errorf("FuzzyDecision() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FuzzyDecision() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHasMultiPrefix(t *testing.T) {
	type args struct {
		code    string
		prefixs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy case",
			args: args{
				code: "SO12345678",
				prefixs: []string{
					"S",
					"O",
				},
			},
			want: true,
		},
		{
			name: "fail case",
			args: args{
				code:    "SO12345678",
				prefixs: []string{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasMultiPrefix(tt.args.code, tt.args.prefixs); got != tt.want {
				t.Errorf("HasMultiPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
