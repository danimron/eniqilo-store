package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, trx domain.Transaction) (domain.Transaction, error)
	// FindByPhoneNumber(ctx context.Context, tx *sql.Tx, phoneNumber string) (domain.Staff, error)
}
