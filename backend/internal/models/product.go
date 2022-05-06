package models

import (
	"mime/multipart"
	"time"
)

type Product struct {
	Id          int                   `json:"id,omitempty"`
	OwnerId     int                   `json:"shop_id,omitempty" validate:"required"`
	Price       float64               `json:"price,omitempty" validate:"required"`
	Name        string                `json:"name,omitempty" validate:"required"`
	Description string                `json:"description,omitempty"`
	Discount    int                   `json:"discount,omitempty"`
	Auction     bool                  `json:"is_auction,omitempty"`
	Selled_at   time.Time             `json:"time,omitempty"`
	Image       *multipart.FileHeader `json:"image"       form:"image"`
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
