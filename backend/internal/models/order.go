package models

import "time"

type Order struct {
	Id            int
	OrderCart     []Product
	OrderedAt     time.Time
	Price         float64
	Discount      *int
	PaymentMethod Payment
}

type Payment struct {
	Digital bool
	COD     bool
}