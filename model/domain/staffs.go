package domain

import (
	"time"
)

type Domain struct {
	Id 			int
	PhoneNumber	string
	Name		string
	Password 	string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt	time.Time
}