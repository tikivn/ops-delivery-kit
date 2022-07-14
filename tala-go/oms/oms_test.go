package oms_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"go.opencensus.io/plugin/ochttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/tikivn/ops-delivery-kit/tala-go/oms"
)

func NewThemisHttp() oms.HttpDoer {
	cfg := clientcredentials.Config{
		ClientID:     "trn:tiki:tms",
		ClientSecret: "m.s0wllvRDCE",
		TokenURL:     "https://api.tala.xyz/oauth2/token",
		Scopes:       []string{"tiki.api"},
		AuthStyle:    oauth2.AuthStyleInParams,
	}
	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{
		Transport: &ochttp.Transport{},
		Timeout:   2 * time.Second,
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	return cfg.Client(ctx)
}

func TestIntegration_client_FetchOrder(t *testing.T) {
	type args struct {
		ctx  context.Context
		code string
	}

	httpDoer := NewThemisHttp()
	c, _ := oms.NewClientWithHost("http://uat.oms.tiki.services", httpDoer)

	pdd := oms.TimeWrap{}
	pdd.UnmarshalJSON(([]byte)(`"2019-03-14 23:59:59"`))
	tests := []struct {
		name    string
		client  oms.Client
		args    args
		want    *oms.Order
		wantErr bool
	}{
		{
			name:   "Real-order",
			client: c,
			args: args{
				ctx:  context.Background(),
				code: "470468764",
			},
			want: &oms.Order{
				OrderID:    8387758,
				Status:     "complete",
				IsRMA:      false,
				Code:       "470468764",
				Substatus:  " ",
				GrandTotal: 22034000,
				Subtotal:   21980000,
				ShippingAddress: oms.ShippingAddress{
					Country:      "Việt Nam",
					WardTikiCode: "VN039022004",
					Ward:         "Phường 04",
					FullName:     "Trần Minh Giàu",
					Phone:        "0336392248",
					Street:       "52 Út Tịch",
					District:     "Quận Tân Bình",
					Region:       "Hồ Chí Minh",
					Email:        "doicanhden@gmail.com",
				},
				ShippingPlan: oms.ShippingPlan{
					PromisedDeliveryDate: pdd,
					PlanID:               1,
					PlanName:             "Giao hàng tiêu chuẩn",
				},
				Payment: oms.Payment{
					Method:    "cod",
					IsPrepaid: false,
				},
				Items: []oms.Item{
					oms.Item{
						ProductID:   309853,
						ProductName: "Samsung Galaxy A7 2017 - Đen",
						ProductSku:  "5803383037169",
						Price:       10990000,
						Qty:         2,
						ProductType: "simple",
					},
				},
				Warehouse: oms.Warehouse{
					WarehouseName: "Hà Nội",
					WarehouseID:   2,
				},
				Shipment: oms.ShipmentWrap{
					Shipment: oms.Shipment{
						PartnerID:    "1",
						PartnerName:  "Tiki Team",
						TrackingCode: "TTEST470468764",
						Status:       "New",
					},
				},
				BackendID: 6685794,
			},
			wantErr: false,
		},
		{
			name:   "Not-exist-order",
			client: c,
			args: args{
				ctx:  context.Background(),
				code: "250208361xxxx",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.client
			got, err := c.FetchOrder(tt.args.ctx, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.FetchOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.FetchOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
