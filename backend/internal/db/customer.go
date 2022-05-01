package db

import (
	"database/sql"
	"secondChance/internal/models"
	"time"
)

type CustomerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (c *CustomerRepo) Get(email string) (*models.Customer, error) {
	var user models.Customer
	sqlStatement := `SELECT customer_id,name,email,password,phone, image FROM customer WHERE email=$1`

	row := c.db.QueryRow(sqlStatement, email)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Image); err != nil {
		return &models.Customer{}, err
	}
	return &user, nil
}

func (c *CustomerRepo) Create(user *models.Customer) error {
	sqlStatement := `INSERT INTO customer (phone, email, name, password) VALUES ($1, $2, $3, $4)`
	if err := c.db.QueryRow(sqlStatement, user.Phone, user.Email, user.Name, user.Password); err != nil {
		return err.Err()
	}
	return nil
}

func (c *CustomerRepo) CreateOrder(order *models.Order) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO order (customer_id, product_id, shop_id) VALUES ($1, $2, $3)`
	if err := tx.QueryRow(sqlStatement, order.Customer_id, order.Shop_id, order.Product_id); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err.Err()

	}

	sqlStatement = `UPDATE product SET selled_at=$2 WHERE product_id=$1`
	if err := tx.QueryRow(sqlStatement, order.Product_id, time.Now()); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err.Err()

	}
	if txErr := tx.Commit(); txErr != nil {
		return txErr
	}
	return nil
}

func (c *CustomerRepo) GetOrder(id *models.IdReg) (*[]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, discount, selled_at from "product" where product_id in (select product_id from "order" where customer_id = $1)`

	rows, err := c.db.Query(sqlStatement, id.Id)
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
