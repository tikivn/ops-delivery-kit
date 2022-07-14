package transaction

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	transaction2 "github.com/tikivn/ops-delivery-kit/transaction"
)

type Tx struct {
	*pg.Tx
}

type transaction struct {
	db *pg.DB
}

type txKeyType struct{}

var txKey txKeyType

func ProvideTransaction(db *pg.DB) transaction2.Transaction {
	return &transaction{db: db}
}

func (t *transaction) Begin(ctx context.Context) (Tx, error) {
	tx, err := t.db.WithContext(ctx).Begin()
	if err != nil {
		return Tx{}, err
	}
	return Tx{
		Tx: tx,
	}, err
}

func ContextWithTransaction(ctx context.Context, tx orm.DB) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func TransactionFromContext(ctx context.Context, fallback orm.DB) orm.DB {
	val := ctx.Value(txKey)
	if tx, ok := val.(orm.DB); ok {
		return tx
	}

	return fallback
}

func (t *transaction) RunWithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.Begin(ctx)
	if err != nil {
		return err
	}

	ctx = ContextWithTransaction(ctx, tx.Tx)
	return tx.RunInTransaction(ctx, func(tx *pg.Tx) error {
		return fn(ctx)
	})
}

func HandleExecuteWithTransactionalInContext(ctx context.Context, db orm.DB, fn func() error) error {
	switch txDB := db.(type) {
	case *pg.Tx:
		return fn()

	case *pg.DB:
		return txDB.RunInTransaction(ctx, func(tx *pg.Tx) error {
			return fn()
		})

	default:
		return errors.New("invalid database instance type to execute transaction")
	}
}
