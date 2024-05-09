package web

import (
	"time"
)

type ProductCreateRes struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
