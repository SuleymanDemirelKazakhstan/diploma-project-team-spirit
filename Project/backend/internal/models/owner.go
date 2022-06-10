package models

import "time"

type Owner struct {
	Id          int    `json:"id"`
	Name        string `json:"name,omitempty" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required,email"`
	Password    string `json:"password,omitempty" validate:"required"`
	Phone       string `json:"phone,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
	Image       string `json:"image"       form:"image"`
	Description string `json:"description,omitempty"`
}

type OwnerFillter struct {
	Id        int       `json:"id,omitempty" validate:"required"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Status    int       `json:"status,omitempty"`
	MinPrice  int       `json:"min_price,omitempty"`
	MaxPrice  int       `json:"max_price,omitempty"`
}

type OwnerProduct struct {
	Id        int       `json:"id,omitempty"`
	Price     float64   `json:"price,omitempty" validate:"required"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Auction   bool      `json:"is_auction,omitempty"`
	Selled_at time.Time `json:"time,omitempty"`
	Status    bool      `json:"status,omitempty"`
	Customer  string    `json:"customer,omitempty"`
}
