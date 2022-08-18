package transaction

import (
	"context"
	"testing"

	"github.com/go-pg/pg/v10"
)

func Test_transaction_RunWithTransaction(t1 *testing.T) {
	type fields struct {
		db *pg.DB
	}
	type args struct {
		ctx context.Context
		fn  func(ctx context.Context) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &transaction{
				db: tt.fields.db,
			}
			if err := t.RunWithTransaction(tt.args.ctx, tt.args.fn); (err != nil) != tt.wantErr {
				t1.Errorf("RunWithTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
