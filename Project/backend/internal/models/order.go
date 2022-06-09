package models

import "time"

type Order struct {
	Customer_id int `json:"customer_id,omitempty" validate:"required"`
	Product_id  int `json:"product_id,omitempty" validate:"required"`
	Shop_id     int `json:"shop_id,omitempty" validate:"required"`
}

type OwnerOrder struct {
	CustomerName string
	ProductName  string
	Price        int
	Auction      bool
	Selled_at    time.Time
	Status       bool
}

type CustomerOrder struct {
	ProductId int
	ShopId    int
	SelledAt  time.Time
	Address   string
	Status    bool
	Image     []string
}
