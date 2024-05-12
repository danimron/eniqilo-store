package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
)

type StaffRepo interface {
	Save(ctx context.Context, tx *sql.Tx, staff domain.Staff) (domain.Staff, error)
	FindByPhoneNumber(ctx context.Context, tx *sql.Tx, phoneNumber string) (domain.Staff, error)
}
