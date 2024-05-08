package domain

import (
	"time"
)

type Customers struct {
	PhoneNumber	string
	Name		string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
}