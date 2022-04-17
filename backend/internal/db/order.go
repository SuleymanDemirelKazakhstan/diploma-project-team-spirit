package db

import "secondChance/internal/models"

func (db *Layer) CreateOrder(order *models.Order) (err error) {
	sqlStatement := `INSERT INTO order (customer_id, product_id, shop_id) VALUES ($1, $2, $3)`
	if err := db.DB.QueryRow(sqlStatement, order.Customer_id, order.Shop_id, order.Product_id); err != nil {
		return err.Err()
	}
	return nil
}
