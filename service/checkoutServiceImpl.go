package service

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
	"eniqilo_store/model/web"
	"eniqilo_store/pkg/errorwrapper"
	"eniqilo_store/repository"

	"github.com/go-playground/validator/v10"
)

type CheckoutServiceImpl struct {
	TransactionRepo repository.TransactionRepository
	ProductRepo     repository.ProductRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func (c CheckoutServiceImpl) Create(ctx context.Context, req web.CheckoutReq) error {
	err := c.Validate.Struct(req)
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusBadRequest, err, "")
	}

	// defer helper.CommitOrRollback(tx)
	products := []domain.Products{}
	totalAmount := 0
	totalItem := 0
	tx, err := c.DB.Begin()
	//process checkout
	err = c.processCheckout(req, &totalAmount, &totalItem, ctx, &products,tx)
	
	if(err != nil){
		return err
	}

	// check amount of paid is correct
	if req.Paid < totalAmount {
		message := "Insufficient amount"
		return errorwrapper.New(errorwrapper.StatusBadRequest, err, message)
	}

	change := req.Paid - totalAmount
	//save transaction
	transaction := domain.Transaction{
		Amount:    req.Paid,
		Change:    change,
		TotalItem: totalItem,
		UserId:    req.CustomerId,
	}
	//save transaction
	_, err = c.TransactionRepo.Save(ctx,tx,transaction)
	if err != nil{
		tx.Rollback()

	}
	//TODO save transaction history
	
	tx.Commit()
	return nil

}

func (c CheckoutServiceImpl) processCheckout(req web.CheckoutReq, totalAmount *int, totalItem *int, ctx context.Context, products *[]domain.Products, tx *sql.Tx) error {
	db := c.DB
	for _, data := range req.ProductDetails {
		product, err := c.ProductRepo.FindById(ctx, db, data.ProductId)
		// check product available
		if err != nil {
			message := "Product Not Found"
			return errorwrapper.New(errorwrapper.StatusNotFound, err, message)
		}
		// check product available
		if !product.IsAvailable {
			message := "Product " + product.Name + " is unavilable"
			return errorwrapper.New(errorwrapper.StatusBadRequest, err, message)
		}
		//check quantity is available
		if product.Stock < data.Quantity {
			message := "Product " + product.Name + " is out of stock"
			return errorwrapper.New(errorwrapper.StatusBadRequest, err, message)
		}
		*products = append(*products, product)
		product.Stock = product.Stock - data.Quantity
		// if stock is empty update product
		if product.Stock == 0 {
			product.IsAvailable = false
		}
		//update stock and available
		_, err = c.ProductRepo.Save(ctx, tx, product)
		if err != nil {
			tx.Rollback()
			return errorwrapper.New(errorwrapper.StatusInternalServerError, err,"")
		}
		*totalAmount += product.Price * data.Quantity
		*totalItem += data.Quantity
		return nil
	}
}
