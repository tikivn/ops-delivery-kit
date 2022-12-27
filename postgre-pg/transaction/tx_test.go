//go:build integration_test
// +build integration_test

package transaction

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transaction_RunWithTransaction(t1 *testing.T) {
	transactional := &transaction{db: db}
	ctx := context.Background()

	t1.Run("Run with external transaction", func(t *testing.T) {
		tx, err := transactional.Begin(ctx)
		assert.Nil(t, err)

		ctx = ContextWithTransaction(ctx, tx)

		_ = transactional.RunWithTransaction(ctx, func(ctx context.Context) error {
			return nil
		})

		assert.False(t, tx.isClosed)
	})
}
