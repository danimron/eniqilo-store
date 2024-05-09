package web

import (
	"time"
)

type ProductStaffGetRes struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Sku         string    `json:"sku"`
	Category    string    `json:"category"`
	imageUrl    string    `json:"imageUrl"`
	Notes       string    `json:"notes"`
	Price       string    `json:"price"`
	Stock       string    `json:"stock"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
}
