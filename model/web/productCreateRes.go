package web

import (
	"time"
)

type ProductCreateRes struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
