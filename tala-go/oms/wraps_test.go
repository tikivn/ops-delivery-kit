package oms

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeWrap_UnmarshalJSON(t *testing.T) {
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
			u := &TimeWrap{
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
