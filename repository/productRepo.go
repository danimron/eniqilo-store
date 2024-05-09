package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	Delete(ctx context.Context, tx *sql.Tx, productId int)
	Update(ctx context.Context, tx *sql.Tx, product domain.Products)
	FindById(ctx context.Context, db *sql.DB, productId int) (domain.Products, error)
	FindAll(ctx context.Context, db *sql.DB, sql string) []domain.Products
}
