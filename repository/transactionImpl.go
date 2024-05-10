package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
	"time"
)

type TransactionRepositoryImpl struct {
}

func NewCheckoutRepository() TransactionRepository {
	return TransactionRepositoryImpl{}
}

func (c TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, trx domain.Transaction) (domain.Transaction, error) {
	sql := "INSERT INTO products(amount, change, total_item, user_id, created_at, updated_at) values($1, $2, $3, $4, $5, $6) RETURNING id"
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, trx.Amount, trx.Change, trx.TotalItem, trx.UserId, time.Now(), time.Now()).Scan(&insertedId)
	trx.Id = insertedId
	return trx, err
}
