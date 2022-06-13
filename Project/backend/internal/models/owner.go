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
	Social      string `json:"social,omitempty"`
}

type OwnerFillter struct {
	Id        int       `json:"id,omitempty" validate:"required"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Status    int       `json:"status,omitempty"`
	MinPrice  int       `json:"min_price,omitempty"`
	MaxPrice  int       `json:"max_price,omitempty"`
	Search    string    `json:"search,omitempty"`
}

type OwnerProduct struct {
	Id        int       `json:"id,omitempty"`
	Price     float64   `json:"price,omitempty" validate:"required"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Auction   bool      `json:"is_auction"`
	Selled_at time.Time `json:"time,omitempty"`
	Status    bool      `json:"status"`
	Customer  string    `json:"customer,omitempty"`
	OrderId   int       `json:"order_id,omitempty"`
}

type CatalogFilter struct {
	Id          int      `json:"id,omitempty" validate:"required"`
	MinPrice    int      `json:"min_price,omitempty"`
	MaxPrice    int      `json:"max_price,omitempty"`
	Category    []string `json:"category,omitempty"`
	Subcategory []string `json:"subcategory,omitempty"`
	Auction     bool     `json:"is_auction,omitempty"`
	Search      string   `json:"search,omitempty"`
}

type DTOowner struct {
	Id      int    `json:"id" validate:"required"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
	Image   string `json:"image"       form:"image"`
	Social  string `json:"social,omitempty"`
}

type MainPage struct {
	Name      string `json:"name,omitempty"`
	Customers int    `json:"customers,omitempty"`
	Earnings  string `json:"earnings,omitempty"`
	Orders    int    `json:"orders,omitempty"`
	Products  int    `json:"products,omitempty"`
}
