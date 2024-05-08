package web

import (
	"time"
)

type ProductsCreateRes struct {
	Id    		string `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
}
