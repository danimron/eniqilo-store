package domain

import (
	"time"
)

type Transaction struct {
	Id        int
	Amount    int
	Change    int
	TotalItem int
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
