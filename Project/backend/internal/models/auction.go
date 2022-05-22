package models

import "time"

type Deal struct {
	ProductId
	Value
}

type Value struct {
	Price      int `json:"price,omitempty"`
	CustomerId int `json:"customer_id,omitempty"`
	StartTime  time.Time
}

type ProductId struct {
	Id string `json:"id,omitempty"`
}
