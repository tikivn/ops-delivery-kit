package publisher

import (
	"context"
	"encoding/json"
	"testing"
)

func TestKafkaPublisher_Publish(t *testing.T) {
	type fields struct {
		brokers []string
		topic   string
		marshal MarshalFunc
	}
	type args struct {
		ctx   context.Context
		key   []byte
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test publish",
			fields: fields{
				brokers: []string{
					"uat-kafka-1.svr.tiki.services:9092",
					"uat-kafka-2.svr.tiki.services:9092",
					"uat-kafka-3.svr.tiki.services:9092",
				},
				topic:   "giau.test-xxxx",
				marshal: json.Marshal,
			},
			args: args{
				ctx: context.Background(),
				key: []byte(`123456`),
				value: map[string]interface{}{
					"orderID": 1231,
					"name":    "Chanh. Ha",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPublisher(tt.fields.brokers, tt.fields.topic)
			if err != nil {
				t.Error(err)
			}
			defer p.Close()

			p.marshal = tt.fields.marshal
			if err := p.Publish(tt.args.ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("KafkaPublisher.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
