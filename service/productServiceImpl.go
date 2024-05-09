package service

import (
	// "cats_social/exception"
	"context"
	"database/sql"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"eniqilo_store/model/web"
	"eniqilo_store/repository"

	// "fmt"
	// "strings"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateReq) web.ProductCreateRes {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product := domain.Products{
		Name:        request.Name,
		Sku:         request.Sku,
		Category:    request.Category,
		ImageUrl:    request.ImageUrl,
		Notes:       request.Notes,
		Price:       request.Price,
		Stock:       request.Stock,
		Location:    request.Location,
		IsAvailable: request.IsAvailable,
	}
	product = service.ProductRepository.Save(ctx, tx, product)
	return helper.ToCategoryResponseCreateProduct(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, ProductId int) {
	db := service.DB
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	_, err = service.ProductRepository.FindById(ctx, db, ProductId)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }
	// if book.Available == 0 {
	// 	message := errors.New("book is booked by someone cannot delete book")
	// 	panic(exception.NewFoundError(message.Error()))
	// }
	service.ProductRepository.Delete(ctx, tx, ProductId)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductCreateReq) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product := domain.Products{
		Name:        request.Name,
		Sku:         request.Sku,
		Category:    request.Category,
		ImageUrl:    request.ImageUrl,
		Notes:       request.Notes,
		Price:       request.Price,
		Stock:       request.Stock,
		Location:    request.Location,
		IsAvailable: request.IsAvailable,
	}
	service.ProductRepository.Update(ctx, tx, product)
}

// func (service *ProductServiceImpl) FindAll(ctx context.Context, cat *web.CatGetParam) []web.CatGetResponse {
// 	db := service.DB
// 	sql := ""

// 	if cat.Id != "" {
// 		sql = AddCondition(sql) + "c.id = " + cat.Id
// 	}
// 	if cat.Race != "" {
// 		sql = AddCondition(sql) + "c.race = '" + cat.Race + "'"
// 	}
// 	if cat.Sex != "" {
// 		sql = AddCondition(sql) + "c.sex = '" + cat.Sex + "'"
// 	}
// 	if cat.Owned != "" {
// 		if strings.ToLower(cat.HasMatched) == "true" {
// 			sql = AddCondition(sql) + "c.user_id IS NOT NULL"
// 		} else {
// 			sql = AddCondition(sql) + "c.user_id IS NULL"
// 		}
// 	}
// 	if cat.HasMatched != "" {
// 		sql = " LEFT JOIN matchs m on (m.issuer_cat_id = c.id or m.receiver_cat_id = c.id) " + sql
// 		if strings.ToLower(cat.HasMatched) == "true" {
// 			sql = AddCondition(sql) + " m.id IS NOT NULL"
// 		} else {
// 			sql = AddCondition(sql) + " m.id IS NULL"
// 		}
// 	}
// 	if cat.AgeInMonth != "" {
// 		if strings.Contains(cat.AgeInMonth, ">") {
// 			sql = AddCondition(sql) + "c.age_in_months " + cat.AgeInMonth
// 		} else if strings.Contains(cat.AgeInMonth, "<") {
// 			sql = AddCondition(sql) + "c.age_in_months " + cat.AgeInMonth
// 		} else {
// 			sql = AddCondition(sql) + "c.age_in_months = " + cat.AgeInMonth
// 		}
// 	}
// 	if cat.Search != "" {
// 		sql = AddCondition(sql) + "c.name LIKE '%" + cat.Search + "%'"
// 	}
// 	if cat.Limit != "" {
// 		sql = sql + " LIMIT " + cat.Limit
// 	}
// 	if cat.Offset != "" {
// 		sql = sql + " OFFSET " + cat.Offset
// 	}
// 	sql = "SELECT c.id, c.name, c.race, c.sex, c.age_in_months, c.description, c.created_at from cats c" + sql
// 	fmt.Println(sql)

// 	cats := service.CatRepository.FindAll(ctx, db, sql)
// 	return helper.ToCategoryResponseCats(cats)
// }

func AddCondition(sql string) string {
	finalSql := ""
	if sql == "" {
		finalSql = sql + " WHERE "
	} else {
		finalSql = sql + " AND "
	}
	return finalSql
}
