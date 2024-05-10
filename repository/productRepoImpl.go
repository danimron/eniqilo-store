package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"time"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Products) (domain.Products, error) {
	sql := "INSERT INTO products(name, sku, category, image_url, notes, price, stock, location, is_available, created_at, updated_at) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, product.Name, product.Sku, product.Category, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, product.IsAvailable, time.Now(), time.Now()).Scan(&insertedId)
	if err != nil {
		return domain.Products{}, err
	}
	product.Id = insertedId
	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int) {
	sql := "UPDATE products SET deleted_at = $1, updated_at = $2 WHERE id = $3"
	_, err := tx.ExecContext(ctx, sql, time.Now(), time.Now(), productId)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Products) {
	sql := "UPDATE products SET name = $1, sku = $2, category = $3, image_url = $4, notes = $5, price = $6, stock = $7, location = $8, is_available = $9, updated_at = $10 WHERE id = $11"
	_, err := tx.ExecContext(ctx, sql, product.Name, product.Sku, product.Category, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, product.IsAvailable, time.Now(), product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, db *sql.DB, productId int) (domain.Products, error) {
	sql := "SELECT * FROM products WHERE id = $1 and deleted_at IS NULL"
	rows, err := db.QueryContext(ctx, sql, productId)
	helper.PanicIfError(err)
	product := domain.Products{}
	defer rows.Close()
	if !rows.Next() {
		err := rows.Scan(&product.Id)
		helper.PanicIfError(err)
		return product, nil
	} else {
		// return cat, errors.New("book is not found")
		return product, nil
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, db *sql.DB, sql string) []domain.Products {
	rows, err := db.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	var products []domain.Products
	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Stock, &product.Notes, &product.Price, &product.Location, &product.IsAvailable, &product.CreatedAt)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}
