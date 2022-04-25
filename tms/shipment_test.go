package tms

import (
	"reflect"
	"testing"
	"time"
)

func TestStTimestamp_MarshalJSON(t1 *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "happy case",
			fields: fields{
				Time: time.Date(2022, 4, 25, 0, 0, 0, 0, time.Local),
			},
			want: []byte(`"2022-04-25 00:00:00"`),
		},
	}
	for _, tt := range tests {
		t1.Run(
			tt.name, func(t1 *testing.T) {
				t := StTimestamp{
					Time: tt.fields.Time,
				}
				got, err := t.MarshalJSON()
				if (err != nil) != tt.wantErr {
					t1.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t1.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
