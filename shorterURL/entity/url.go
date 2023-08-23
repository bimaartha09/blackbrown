package entity

import (
	"time"
)

type Url struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Destination string    `json:"destination"`
	Counter     int64     `json:"counter"`
	ExpireTime  time.Time `json:"expire_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var Urls = []Url{}
