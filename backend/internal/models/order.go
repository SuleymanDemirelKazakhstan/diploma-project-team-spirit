package models

type Order struct {
	Customer_id int `json:"customer_id,omitempty" validate:"required"`
	Product_id  int `json:"product_id,omitempty" validate:"required"`
	Shop_id     int `json:"shop_id,omitempty" validate:"required"`
}
