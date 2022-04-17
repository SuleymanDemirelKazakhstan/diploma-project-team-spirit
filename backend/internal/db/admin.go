package db

import (
	"secondChance/internal/models"
)

func (db *Layer) CreateOwner(user *models.Owner) error {
	sqlStatement := `INSERT INTO owner (phone, email, name, password, address) VALUES ($1, $2, $3, $4)`
	if err := db.DB.QueryRow(sqlStatement, user.Phone, user.Email, user.Name, user.Password, user.Address); err != nil {
		return err.Err()
	}
	return nil
}

func (db *Layer) GetOwner(param *models.OwnerEmailRequest) (*models.Owner, error) {
	var user models.Owner
	sqlStatement := `SELECT shop_id,name,email,password,phone,address FROM owner WHERE email=$1`

	row := db.DB.QueryRow(sqlStatement, param.Email)
	// unmarshal the row object to user
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Address); err != nil {
		return &models.Owner{}, err
	}
	return &user, nil
}

func (db *Layer) DeleteOwner(param *models.OwnerEmailRequest) error {
	sqlStatement := `DELETE FROM owner WHERE email=$1`
	if _, err := db.DB.Exec(sqlStatement, param.Email); err != nil {
		return err
	}
	return nil
}

func (db *Layer) GetAllOwner() ([]models.Owner, error) {
	var users []models.Owner
	sqlStatement := `SELECT name,email,password,phone,address FROM owner`

	rows, err := db.DB.Query(sqlStatement)
	if err != nil {
		return []models.Owner{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.Owner
		if err := rows.Scan(&user.Name, &user.Email, &user.Password, &user.Phone, &user.Address); err != nil {
			return []models.Owner{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *Layer) UpdateOwner(email *models.OwnerEmailRequest, user *models.Owner) error {
	sqlStatement := `UPDATE owner SET name=$2, password=$3, phone=$4, email=$5, address=$6 WHERE email=$1`
	_, err := db.DB.Exec(sqlStatement, email.Email, user.Name, user.Password, user.Phone, user.Email, user.Address)
	if err != nil {
		return err
	}
	return nil
}
