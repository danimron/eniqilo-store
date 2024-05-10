package domain

import (
	"time"
)

type TransactionDetail struct {
	Id            int
	TransactionId int
	ProductId     int
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
