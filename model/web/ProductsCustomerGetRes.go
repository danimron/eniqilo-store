package web

import (
	"time"
)

type ProductsCustomerRes struct {
	Id    		string `json:"id"`
	Name    	string `json:"name"`
	Sku 		string `json:"sku"`
	Category    string `json:"category"`
	imageUrl    string `json:"imageUrl"`
	Price     	string `json:"price"`
	Stock		string `json:"stock"`
	Location    string `json:"location"`
	CreatedAt   time.Time `json:"createdAt"`
}
