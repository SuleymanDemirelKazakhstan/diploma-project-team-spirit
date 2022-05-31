package models

import (
	"time"
)

type Product struct {
	Id          int       `json:"id,omitempty"`
	OwnerId     int       `json:"shop_id,omitempty" validate:"required"`
	Price       float64   `json:"price,omitempty" validate:"required"`
	Name        string    `json:"name,omitempty" validate:"required"`
	Description string    `json:"description,omitempty" validate:"required"`
	Discount    int       `json:"discount,omitempty"`
	Auction     bool      `json:"is_auction,omitempty"`
	Category    string    `json:"category,omitempty" validate:"required"`
	Subcategory string    `json:"subcategory,omitempty" validate:"required"`
	Size        string    `json:"size,omitempty" validate:"required"`
	Colour      string    `json:"colour,omitempty" validate:"required"`
	Condition   string    `json:"condition,omitempty" validate:"required"`
	Selled_at   time.Time `json:"time,omitempty"`
	Image       []string  `json:"image"       form:"image"`
}

type Filter struct {
	Category    string `json:"category,omitempty"`
	Subcategory string `json:"subcategory,omitempty"`
	Size        string `json:"size,omitempty"`
	Colour      string `json:"colour,omitempty"`
	Condition   string `json:"condition,omitempty"`
}

type IdReg struct {
	Id int `json:"id,omitempty" validate:"required"`
}
