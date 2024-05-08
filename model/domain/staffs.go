package domain

import (
	"time"
)

type Staffs struct {
	Id 			int
	PhoneNumber	string
	Name		string
	Password 	string
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt	time.Time
}