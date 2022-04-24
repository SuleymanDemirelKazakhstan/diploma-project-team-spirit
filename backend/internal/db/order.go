package db

import "secondChance/internal/models"

func (db *Layer) CreateOrder(order *models.Order) (err error) {
	sqlStatement := `INSERT INTO order (customer_id, product_id, shop_id) VALUES ($1, $2, $3)`
	if err := db.DB.QueryRow(sqlStatement, order.Customer_id, order.Shop_id, order.Product_id); err != nil {
		return err.Err()
	}
	return nil
}

func (db *Layer) CustomerOrder(id *models.IdReg) (*[]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, discount, selled_at from "product" where product_id in (select product_id from "order" where customer_id = $1)`

	rows, err := db.DB.Query(sqlStatement, id.Id)
	if err != nil {
		return &[]models.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Price, &product.Name, &product.Discount, &product.Selled_at); err != nil {
			return &[]models.Product{}, err
		}
		products = append(products, product)
	}
	return &products, nil
}

func (db *Layer) OwnerOrder(id *models.IdReg) (*[]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, discount, selled_at from "product" where product_id in (select product_id from "order" where shop_id = $1)`

	rows, err := db.DB.Query(sqlStatement, id.Id)
	if err != nil {
		return &[]models.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Price, &product.Name, &product.Discount, &product.Selled_at); err != nil {
			return &[]models.Product{}, err
		}
		products = append(products, product)
	}
	return &products, nil
}
