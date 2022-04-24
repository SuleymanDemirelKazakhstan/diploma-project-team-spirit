package models

import "time"

type Product struct {
	Id          int     `json:"id,omitempty"`
	OwnerId     int     `json:"owner_id,omitempty" validate:"required"`
	Price       float64 `json:"price,omitempty" validate:"required"`
	Name        string  `json:"name,omitempty" validate:"required"`
	Description string  `json:"description,omitempty"`
	Discount    int     `json:"discount,omitempty"`
	IsAuction
	Selled_at time.Time `json:"time,omitempty"`
}

type Products struct {
	Id      int     `json:"id,omitempty"`
	OwnerId int     `json:"owner_id,omitempty"`
	Price   float64 `json:"price,omitempty"`
	Name    string  `json:"name,omitempty"`
}

type IdReg struct {
	Id int `json:"id,omitempty" validate:"required"`
}

type IsAuction struct {
	Auction bool `json:"is_auction,omitempty" validate:"required"`
}
