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

func TestRound(t *testing.T) {
	type args struct {
		val       float64
		roundOn   float64
		precision int
	}
	tests := []struct {
		name       string
		args       args
		wantNewVal float64
	}{
		{
			name: "Over point",
			args: args{
				val:       10.662435,
				roundOn:   0.5,
				precision: 1,
			},
			wantNewVal: 10.7,
		},
		{
			name: "Not over point",
			args: args{
				val:       10.662435,
				roundOn:   0.5,
				precision: 2,
			},
			wantNewVal: 10.66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewVal := Round(tt.args.val, tt.args.roundOn, tt.args.precision); gotNewVal != tt.wantNewVal {
				t.Errorf("Round() = %v, want %v", gotNewVal, tt.wantNewVal)
			}
		})
	}
}
