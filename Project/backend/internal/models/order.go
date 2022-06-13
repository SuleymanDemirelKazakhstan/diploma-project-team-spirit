package models

import "time"

type Order struct {
	Customer_id int `json:"customer_id,omitempty" validate:"required"`
	Product_id  int `json:"product_id,omitempty" validate:"required"`
	Shop_id     int `json:"shop_id,omitempty" validate:"required"`
}

type OwnerOrder struct {
	CustomerName  string `json:"customer_name,omitempty"`
	CustomerEmail string `json:"customer_email,omitempty"`
	ProductId     int    `json:"product_id,omitempty"`
	ProductName   string `json:"product_name,omitempty"`
	Size          string `json:"siize,omitempty"`
	Price         string `json:"price,omitempty"`
	Image         []string
	Create_at     time.Time `json:"create_time,omitempty"`
	Selled_at     time.Time `json:"selled_time,omitempty"`
	Status        bool      `json:"status"`
}

type CustomerOrder struct {
	ProductId int
	ShopId    int
	SelledAt  time.Time
	Address   string
	Status    bool
	Image     []string
}
