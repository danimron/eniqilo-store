package domain

import (
	"time"
)

type Session struct {
	Id        int
	StaffId   int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
