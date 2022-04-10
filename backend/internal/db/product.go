package db

import (
	"secondChance/internal/models"
	"time"
)

func (db *Layer) CreateProduct(product *models.Product) (err error){
	sqlStatement := `INSERT INTO product (owner_id, price, name, description) VALUES ($1, $2, $3, $4)`
	if err := db.DB.QueryRow(sqlStatement, product.OwnerId, product.Price, product.Name, product.Description); err != nil {
		return err.Err()
	}
	return nil
}

func (db *Layer) GetProduct(id *models.IdReg) (*models.Product, error){
	var product models.Product
	sqlStatement := `SELECT owner_id,price,name,description FROM product WHERE id=$1`

	row := db.DB.QueryRow(sqlStatement, id.Id)
	// unmarshal the row object to user
	if err := row.Scan(&product.OwnerId, &product.Price, &product.Name, &product.Description); err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}

func (db *Layer) GetAllProduct() ([]models.Products, error){
	var products []models.Products
	sqlStatement := `SELECT id, owner_id, price, name from product where selled_at is null`

	rows, err := db.DB.Query(sqlStatement)
	if err != nil {
		return []models.Products{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Products
		if err := rows.Scan(&product.Id, &product.OwnerId, &product.Price, &product.Name); err != nil {
			return []models.Products{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (db *Layer) SoldProduct(t time.Time, id models.IdReg) error{
	sqlStatement := `UPDATE product SET selled_at=$2 WHERE id=$1`
	if err := db.DB.QueryRow(sqlStatement, id, t); err != nil {
		return err.Err()
	}
	return nil
}

func (db *Layer) UpdateProduct(id *models.IdReg,product *models.Product)  error{
	sqlStatement := `UPDATE product SET price=$2, name=$3, description=$4 WHERE id=$1`
	_, err := db.DB.Exec(sqlStatement, id.Id, product.Price, product.Name, product.Description)
	if err != nil {
		return err
	}
	return nil
}
