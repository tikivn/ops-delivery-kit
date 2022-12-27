package transaction

import (
	"context"
	"reflect"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"

	transaction2 "github.com/tikivn/ops-delivery-kit/transaction"
)

type Tx struct {
	*pg.Tx
	isClosed bool
}

func (tx Tx) RunTransaction(ctx context.Context, fn func(tx *pg.Tx) error) error {
	tx.isClosed = true
	return tx.Tx.RunInTransaction(ctx, fn)
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
	tx := TransactionFromContext(ctx, nil)

	// Nếu transaction chưa được khởi tạo thì thời điểm hiện tại là thời điểm transaction bắt đầu
	if tx == nil {
		tr, err := t.Begin(ctx)
		if err != nil {
			return err
		}

		ctx = ContextWithTransaction(ctx, tr)

		// thực hiện commit cho transaction
		return tr.RunTransaction(ctx, func(tx *pg.Tx) error {
			return fn(ctx)
		})
	}

	// Transaction đã được khởi tạo trước đó, để cho layer phía trên tự commit
	return fn(ctx)
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
		txDbType := reflect.TypeOf(txDB)
		return errors.Errorf("invalid database instance type to execute transaction with type: %s", txDbType.Name())
	}
}
