package service

import (
	// "cats_social/exception"
	"context"
	"database/sql"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"eniqilo_store/model/web"
	"eniqilo_store/pkg/errorwrapper"
	"eniqilo_store/repository"
	"net/url"
	"slices"

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

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateReq) (web.ProductCreateRes, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.ProductCreateRes{}, errorwrapper.New(errorwrapper.StatusBadRequest, err, "")
	}

	tx, err := service.DB.Begin() // transaction db
	if err != nil {
		return web.ProductCreateRes{}, errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}

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

	product, err = service.ProductRepository.Save(ctx, tx, product)
	errMsg := validateRequestValue(product)
	if errMsg != "" {
		return web.ProductCreateRes{}, errorwrapper.New(errorwrapper.StatusBadRequest, err, "")
	}

	if err != nil {
		tx.Rollback()
		return web.ProductCreateRes{}, errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}

	return helper.ToCategoryResponseCreateProduct(product), nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, ProductId int) error {
	db := service.DB
	tx, err := service.DB.Begin()
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}

	_, err = service.ProductRepository.FindById(ctx, db, ProductId)
	if err != nil {
		message := "Id is Not Found"
		tx.Rollback()
		return errorwrapper.New(errorwrapper.StatusNotFound, err, message)
	}

	err = service.ProductRepository.Delete(ctx, tx, ProductId)
	if err != nil {
		tx.Rollback()
		return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}
	return nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductCreateReq) error {
	err := service.Validate.Struct(request)
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusBadRequest, err, "")
	}

	tx, err := service.DB.Begin() // transaction db
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}

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

	errMsg := validateRequestValue(product)
	if errMsg != "" {
		return errorwrapper.New(errorwrapper.StatusBadRequest, err, "")
	}

	err = service.ProductRepository.Update(ctx, tx, product)
	if err != nil {
		tx.Rollback()
		return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}
	return nil
}

func validateRequestValue(product domain.Products) string {
	category := []string{"Clothing", "Accessories", "Footwear", "Beverages"}
	if len(product.Name) < 1 || len(product.Name) > 30 {
		return "Name length must be in range 1 - 30 character"
	}
	if len(product.Sku) < 1 || len(product.Sku) > 30 {
		return "SKU length must be in range 1 - 30 character"
	}
	if slices.Contains(category, product.Category) {
		return "Category not in list"
	}
	_, err := url.ParseRequestURI(product.ImageUrl)
	if err != nil {
		return "Image Url format is not valid"
	}
	if len(product.Notes) < 1 || len(product.Notes) > 200 {
		return "Notes length must be in range 1 - 200 character"
	}
	if product.Price == 0 || product.Price > 100000 {
		return "Price must be in range 1 - 100000"
	}
	if len(product.Location) < 1 || len(product.Location) > 200 {
		return "Notes length must be in range 1 - 200 character"
	}
	if !product.IsAvailable == true || !product.IsAvailable == false {
		return "IsAvailable type must be boolean"
	}

	return ""
}
