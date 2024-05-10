package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
	"time"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return TransactionRepositoryImpl{}
}

func (c TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, trx domain.Transaction) (domain.Transaction, error) {
	sql := "INSERT INTO transactions(amount, change, total_item, user_id, created_at, updated_at) values($1, $2, $3, $4, $5, $6) RETURNING id"
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, trx.Amount, trx.Change, trx.TotalItem, trx.UserId, time.Now(), time.Now()).Scan(&insertedId)
	trx.Id = insertedId
	return trx, err
}

func (c TransactionRepositoryImpl) FindTransactionHistory(ctx context.Context, sql string, db *sql.DB, trxId int, arg ...any) ([]domain.TransactionHistory, error) {
	rows, err := db.QueryContext(ctx, sql, arg...)
	if err != nil {
		return []domain.TransactionHistory{}, err
	}
	trxHistories := []domain.TransactionHistory{}
	defer rows.Close()
	if !rows.Next() {
		trxHistory := domain.TransactionHistory{}
		productDetail := domain.ProductDetails{}
		err := rows.Scan(&trxHistory.TransactionId, &trxHistory.CustomerId, &productDetail, &trxHistory.Paid, &trxHistory.Change, &trxHistory.CreatedAt)
		if err != nil {
			return []domain.TransactionHistory{}, err
		}
		trxHistory.ProductDetails = append(trxHistory.ProductDetails, productDetail)
		trxHistories = append(trxHistories, trxHistory)
		return trxHistories, nil
	} else {
		// return cat, errors.New("book is not found")
		return trxHistories, nil
	}
}
