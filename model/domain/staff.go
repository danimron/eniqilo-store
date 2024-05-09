package domain

import (
	"time"
)

type Staff struct {
	Id          int
	PhoneNumber string
	Name        string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
