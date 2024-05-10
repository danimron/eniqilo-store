package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"time"
)

type TransactionDetailRepositoryImpl struct {
}

func NewTransactionDetailRepository() TransactionDetailRepository {
	return TransactionDetailRepositoryImpl{}
}
func (td TransactionDetailRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, trx domain.TransactionDetail) (domain.TransactionDetail, error) {
	sql := "INSERT INTO products(transaction_id, product_id, quantity,created_at, updated_at) values($1, $2, $3, $4, $5) RETURNING id"
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, trx.TransactionId, trx.ProductId, trx.Quantity, time.Now(), time.Now()).Scan(&insertedId)
	trx.Id = insertedId
	return trx, err
	// panic("not implemented") // TODO: Implement
}

func (td TransactionDetailRepositoryImpl) FindByTransactionId(ctx context.Context, db *sql.DB, id int) (domain.TransactionDetail, error) {
	sql := "SELECT id, transacion_id, product_id, quantity, created_at FROM transaction_detail WHERE transaction_id = $1"
	rows, err := db.QueryContext(ctx, sql, id)
	if err != nil{
		return domain.TransactionDetail{}, err
	}
	trxDetail := domain.TransactionDetail{}
	defer rows.Close()
	if !rows.Next() {
		err := rows.Scan(&trxDetail.Id, &trxDetail.TransactionId, &trxDetail.ProductId, &trxDetail.Quantity, &trxDetail.CreatedAt)
		helper.PanicIfError(err)
		return trxDetail, nil
	} else {
		// return cat, errors.New("book is not found")
		return trxDetail, nil
	}
}
