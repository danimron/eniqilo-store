package domain

import (
	"time"
)

type Domain struct {
	Id 			int
	Name		string
	Sku		 	string
	Category	string
	ImageUrl	string
	Notes		string
	Price		int
	Stock		int
	Location 	string
	IsAvailable	bool
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt	time.Time
}