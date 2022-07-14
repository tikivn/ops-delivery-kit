package algorithm

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/tikivn/ops-delivery-kit/util"
)

func TestRemoveSliceKey(t *testing.T) {
	type args struct {
		sl          []removeable
		removeables []removeableKey
	}
	tests := []struct {
		name string
		args args
		want []removeable
	}{
		{
			name: "Test remove",
			args: args{
				sl: []removeable{
					assignBox{
						BoxID: util.CodeToUUID("1"),
					},
					assignBox{
						BoxID: util.CodeToUUID("2"),
					},
					assignBox{
						BoxID: util.CodeToUUID("3"),
					},
					assignBox{
						BoxID: util.CodeToUUID("4"),
					},
					assignBox{
						BoxID: util.CodeToUUID("5"),
					},
				},
				removeables: []removeableKey{
					util.CodeToUUID("1"),
					util.CodeToUUID("4"),
					util.CodeToUUID("3"),
				},
			},
			want: []removeable{
				assignBox{
					BoxID: util.CodeToUUID("2"),
				},
				assignBox{
					BoxID: util.CodeToUUID("5"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSliceKey(tt.args.sl, tt.args.removeables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSliceKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

type assignBox struct {
	BoxID uuid.UUID
}

func (a assignBox) Key() removeableKey {
	return a.BoxID
}
