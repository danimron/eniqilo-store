package domain

import (
	"time"
)

type Sessions struct {
	Id        int
	Token     string
	CreatedAt time.Time
}
