package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
)

type SessionRepo interface {
	Save(ctx context.Context, tx *sql.Tx, staff domain.Staff, token string) error
	FindByStaffId(ctx context.Context, tx *sql.Tx, staffId int) error
}
