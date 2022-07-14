package util

import (
	"testing"
)

func Test_mapField(t *testing.T) {
	type args struct {
		obj   interface{}
		tag   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mapField(tt.args.obj, tt.args.tag, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("mapField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
