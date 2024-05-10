package web

import "eniqilo_store/model/domain"

type CheckoutHistoryResponse struct {
	TransactionId  int
	CustomerId     int
	ProductDetails []domain.ProductDetails
	Paid           int
	Change         int
}
