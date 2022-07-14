package util

import (
	"testing"
	"time"
)

func TestSubtractByEpochTime(t *testing.T) {
	currentTimeUTC := time.Date(2021, time.July, 29, 11, 0, 0, 0, time.UTC)

	hcmZone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	currentTimeWithTimezone := time.Date(2021, time.July, 29, 18, 0, 0, 0, hcmZone)

	type args struct {
		currentTime time.Time
		epochTime   int64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "happy case",
			args: args{
				currentTime: currentTimeUTC,
				epochTime:   1627557180,
			},
			want: time.Duration(780) * time.Second,
		},
		{
			name: "wrong timezone",
			args: args{
				currentTime: currentTimeWithTimezone,
				epochTime:   1627557180,
			},
			want: time.Duration(780) * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubtractByEpochTime(tt.args.epochTime, tt.args.currentTime); got != tt.want {
				t.Errorf("SubtractByEpochTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
