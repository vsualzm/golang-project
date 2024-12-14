package models

import "time"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
	CreateAt time.Time
	UpdateAt time.Time
}

type ProductInputCreate struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
