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
	TransactionRepo       repository.TransactionRepository
	TransactionDetailRepo repository.TransactionDetailRepository
	ProductRepo           repository.ProductRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewCheckoutService(checkoutService CheckoutService, 
	transactionRepo repository.TransactionRepository,
	 trxDetailRepo repository.TransactionDetailRepository,
	 productRepo repository.ProductRepository,
	 db *sql.DB, validate *validator.Validate ) CheckoutService{
	return CheckoutServiceImpl{
		TransactionRepo: transactionRepo,
		TransactionDetailRepo: trxDetailRepo,
		ProductRepo: productRepo,
		DB: db,
		Validate: validate,
	}
}

func (c CheckoutServiceImpl) Create(ctx context.Context, req web.CheckoutReq) error {
	err := c.Validate.Struct(req)
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusBadRequest, err, "")
	}

	// defer helper.CommitOrRollback(tx)
	trxDetails := []domain.TransactionDetail{}
	totalAmount := 0
	totalItem := 0
	tx, err := c.DB.Begin()
	if err != nil {
		return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
	}
	//process checkout
	err = c.processCheckout(req, &totalAmount, &totalItem, ctx, &trxDetails, tx)

	if err != nil {
		return err
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
	trx, err := c.TransactionRepo.Save(ctx, tx, transaction)

	if err != nil {
		tx.Rollback()

	}
	//TODO save transaction Detail
	for _, trxDetail := range trxDetails {
		trxDetail.TransactionId = trx.Id
		_, err = c.TransactionDetailRepo.Save(ctx, tx, trxDetail)
		if err != nil {
			return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		}
	}

	tx.Commit()
	return nil

}

func (c CheckoutServiceImpl) processCheckout(req web.CheckoutReq, totalAmount *int, totalItem *int, ctx context.Context, trxDetails *[]domain.TransactionDetail, tx *sql.Tx) error {
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
			return errorwrapper.New(errorwrapper.StatusBadRequest, nil, message)
		}
		//check quantity is available
		if product.Stock < data.Quantity {
			message := "Product " + product.Name + " is out of stock"
			return errorwrapper.New(errorwrapper.StatusBadRequest, nil, message)
		}
		trxDetail := domain.TransactionDetail{
			ProductId: product.Id,
			Quantity:  data.Quantity,
		}
		*trxDetails = append(*trxDetails, trxDetail)
		product.Stock = product.Stock - data.Quantity
		// if stock is empty update product
		if product.Stock == 0 {
			product.IsAvailable = false
		}
		//update stock and available
		_, err = c.ProductRepo.Save(ctx, tx, product)
		if err != nil {
			tx.Rollback()
			return errorwrapper.New(errorwrapper.StatusInternalServerError, err, "")
		}
		*totalAmount += product.Price * data.Quantity
		*totalItem += data.Quantity
	}

	// check amount of paid is correct
	if req.Paid < *totalAmount {
		message := "Insufficient amount"
		return errorwrapper.New(errorwrapper.StatusBadRequest, nil, message)
	}

	return nil
}
