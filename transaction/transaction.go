package transaction

import (
	"context"
)

type Transaction interface {
	RunWithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
