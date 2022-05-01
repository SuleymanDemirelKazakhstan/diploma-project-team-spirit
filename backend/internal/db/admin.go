package db

import (
	"database/sql"
	"secondChance/internal/models"
)

type AdminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

func (a *AdminRepo) Create(user *models.Owner) error {
	sqlStatement := `INSERT INTO shop (phone, email, name, password, address) VALUES ($1, $2, $3, $4, $5)`
	if err := a.db.QueryRow(sqlStatement, user.Phone, user.Email, user.Name, user.Password, user.Address); err != nil {
		return err.Err()
	}
	return nil
}

func (a *AdminRepo) Get(param *models.OwnerEmailRequest) (*models.Owner, error) {
	var user models.Owner
	sqlStatement := `SELECT shop_id,name,email,password,phone,address,image FROM shop WHERE email=$1 and is_deleted=false`

	row := a.db.QueryRow(sqlStatement, param.Email)
	// unmarshal the row object to user
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Address, &user.Image); err != nil {
		return &models.Owner{}, err
	}
	return &user, nil
}

func (a *AdminRepo) Delete(param *models.OwnerEmailRequest) error {
	sqlStatement := `update shop set is_deleted=true WHERE email=$1`
	if _, err := a.db.Exec(sqlStatement, param.Email); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepo) GetAll() ([]models.Owner, error) {
	var users []models.Owner
	sqlStatement := `SELECT name,email,password,phone,address FROM owner where is_deleted=false`

	rows, err := a.db.Query(sqlStatement)
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

func (a *AdminRepo) Update(email *models.OwnerEmailRequest, user *models.Owner) error {
	sqlStatement := `UPDATE shop SET name=$2, password=$3, phone=$4, email=$5, address=$6 WHERE email=$1`
	_, err := a.db.Exec(sqlStatement, email.Email, user.Name, user.Password, user.Phone, user.Email, user.Address)
	if err != nil {
		return err
	}
	return nil
}
