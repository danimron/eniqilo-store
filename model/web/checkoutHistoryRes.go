package web

type CheckoutHistoryResponse struct {
	TransactionId  int
	CustomerId     int
	ProductDetails []ProductDetails
	Paid           int
	Change         int
}
