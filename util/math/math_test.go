package math

import "testing"

func TestAbs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test happy case negative number",
			args: args{
				n: -5,
			},
			want: 5,
		},
		{
			name: "Test happy case positive number",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.n); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundToBreakpoint(t *testing.T) {
	type args struct {
		val        float64
		roundPoint float64
		precision  uint
	}
	tests := []struct {
		name       string
		args       args
		wantNewVal float64
		wantErr    bool
	}{
		{
			name: "not over point",
			args: args{
				val:        10.66345,
				roundPoint: 0.5,
				precision:  2,
			},
			wantNewVal: 10.66,
			wantErr:    false,
		},
		{
			name: "over point",
			args: args{
				val:        10.66345,
				roundPoint: 0.5,
				precision:  1,
			},
			wantNewVal: 10.7,
			wantErr:    false,
		},
		{
			name: "invalid point",
			args: args{
				val:        10.66345,
				roundPoint: 1,
				precision:  2,
			},
			wantNewVal: 0,
			wantErr:    true,
		},
		{
			name: "negative precision",
			args: args{
				val:        -10.66345,
				roundPoint: 0.5,
				precision:  1,
			},
			wantNewVal: -10.7,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewVal, err := RoundToBreakpoint(tt.args.val, tt.args.roundPoint, tt.args.precision)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoundToBreakpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNewVal != tt.wantNewVal {
				t.Errorf("RoundToBreakpoint() gotNewVal = %v, want %v", gotNewVal, tt.wantNewVal)
			}
		})
	}
}
