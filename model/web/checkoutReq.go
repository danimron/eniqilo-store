package web

type CheckoutReq struct {
	CustomerId     int
	ProductDetails []ProductDetails
	Paid           int
	Change         int
}
