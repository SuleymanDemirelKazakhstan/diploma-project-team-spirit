package db

import "secondChance/internal/models"

func (db *Layer) GetCustomer(param string) (*models.Customer, error) {
	var user models.Customer
	sqlStatement := `SELECT id,name,email,password,phone FROM customer WHERE email=$1`

	row := db.DB.QueryRow(sqlStatement, param)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone); err != nil {
		return &models.Customer{}, err
	}
	return &user, nil
}

func (db *Layer) CreateCustomer(user *models.Customer) error {
	sqlStatement := `INSERT INTO customer (phone, email, name, password) VALUES ($1, $2, $3, $4)`
	if err := db.DB.QueryRow(sqlStatement, user.Phone, user.Email, user.Name, user.Password); err != nil {
		return err.Err()
	}
	return nil
}
