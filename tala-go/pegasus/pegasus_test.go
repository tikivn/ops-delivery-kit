package pegasus

import (
	"context"
	"fmt"
	"testing"
)

func Test_client_QueryProducts(t *testing.T) {
	type args struct {
		ctx        context.Context
		productIDs []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test list",
			args: args{
				ctx:        context.Background(),
				productIDs: []int64{10579998, 8139369, 6701827, 641794},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(HostQueryStaging, HostSingleStaging, nil)
			got, err := c.QueryProducts(tt.args.ctx, tt.args.productIDs)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("client.QueryProducts() = %v\n", got)
		})
	}
}
