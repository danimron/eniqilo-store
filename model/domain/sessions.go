package domain

import (
	"time"
)

type Domain struct {
	Id        int
	Token     string
	CreatedAt time.Time
}
