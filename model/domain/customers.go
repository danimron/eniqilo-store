package domain

import (
	"time"
)

type Domain struct {
	PhoneNumber	string
	Name		string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
}