package domain

import (
	"time"
)

type Customers struct {
	Id          int
	PhoneNumber string
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
