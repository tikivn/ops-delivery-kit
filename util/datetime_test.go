package util

import (
	"reflect"
	"testing"
	"time"
)

func TestDateTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    time.Time
		wantErr bool
	}{
		//
		{
			name:   "Happy case",
			fields: fields{},
			args: args{
				data: []byte(`"2018-10-15 21:00:00"`),
			},
			want: time.Date(2018, time.October, 15, 21, 0, 0, 0, time.Local),
		},
		{
			name:   "Happy case (null)",
			fields: fields{},
			args: args{
				data: []byte(`null`),
			},
		},
		{
			name:   "php empty string",
			fields: fields{},
			args: args{
				data: []byte(`""`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &DateTime{
				Time: tt.fields.Time,
			}
			if err := u.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := u.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(u.Time, tt.want) {
				t.Error(
					"For input: ", string(tt.args.data),
					"expected:", tt.want,
					"got:", u)
			}
		})
	}
}

func TestParseStringsToTime(t *testing.T) {
	vnLocation, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	type args struct {
		dateString  string
		timeString  string
		timeZone    string
		date_layout string
		time_layout string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				dateString:  "2022-04-29",
				timeString:  "09:00",
				timeZone:    "Asia/Ho_Chi_Minh",
				date_layout: "2006-01-02",
				time_layout: "15:04",
			},
			want: time.Date(2022, 4, 29, 9, 0, 0, 0, vnLocation),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringsToTime(tt.args.dateString, tt.args.timeString, tt.args.timeZone, tt.args.date_layout, tt.args.time_layout)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringsToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringsToTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
