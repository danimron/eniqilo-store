package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
)

type TransactionDetailRepository interface {
	Save(ctx context.Context, tx *sql.Tx, trx domain.TransactionDetail) (domain.TransactionDetail, error)
	FindByTransactionId(ctx context.Context, db *sql.DB, id int) (domain.TransactionDetail, error)
}
