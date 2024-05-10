package web

import "eniqilo_store/model/domain"

type CheckoutReq struct {
	CustomerId     int
	ProductDetails []domain.ProductDetails
	Paid           int
	Change         int
}
