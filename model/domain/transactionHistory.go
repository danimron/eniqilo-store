package domain

import "time"

type TransactionHistory struct {
	TransactionId  int
	CustomerId     int
	ProductDetails []ProductDetails
	Paid           int
	Change         int
	CreatedAt      time.Time
}
