package models

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}