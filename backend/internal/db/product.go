package db

import (
	"secondChance/internal/models"
	"time"
)

func (db *Layer) CreateProduct(product *models.Product) (err error) {
	sqlStatement := `INSERT INTO product (shop_id, price, name, description, is_auction) VALUES ($1, $2, $3, $4)`
	if err := db.DB.QueryRow(sqlStatement, product.OwnerId, product.Price, product.Name, product.Description); err != nil {
		return err.Err()
	}
	return nil
}

func (db *Layer) GetProduct(id *models.IdReg) (*models.Product, error) {
	var product models.Product
	sqlStatement := `SELECT shop_id,price, name,description, discount FROM product WHERE product_id=$1`

	row := db.DB.QueryRow(sqlStatement, id.Id)
	// unmarshal the row object to user
	if err := row.Scan(&product.OwnerId, &product.Price, &product.Name, &product.Description, &product.Discount); err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}

func (db *Layer) GetProductAuction(id *models.IdReg) (*models.Product, error) {
	var product models.Product
	sqlStatement := `SELECT shop_id,price, name,description FROM product WHERE product_id=$1 and is_auction=true`
	row := db.DB.QueryRow(sqlStatement, id.Id)
	// unmarshal the row object to user
	if err := row.Scan(&product.OwnerId, &product.Price, &product.Name, &product.Description); err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}

func (db *Layer) GetAllProduct(b *models.IsAuction) ([]models.Products, error) {
	var products []models.Products
	sqlStatement := `SELECT product_id, shop_id, price, name from product where selled_at is null and is_auction=$1`

	rows, err := db.DB.Query(sqlStatement, b.Auction)
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

func (db *Layer) SoldProduct(t time.Time, id int) error {
	sqlStatement := `UPDATE product SET selled_at=$2 WHERE product_id=$1`
	if err := db.DB.QueryRow(sqlStatement, id, t); err != nil {
		return err.Err()
	}
	return nil
}

func (db *Layer) UpdateProduct(id *models.IdReg, product *models.Product) error {
	sqlStatement := `UPDATE product SET price=$2, name=$3, description=$4, discount=$5, is_auction=$6 WHERE id=$1`
	_, err := db.DB.Exec(sqlStatement, id.Id, product.Price, product.Name,
		product.Description, product.Discount, product.IsAuction.Auction)
	if err != nil {
		return err
	}
	return nil
}

func (db *Layer) DeleteProduct(param *models.IdReg) error {
	sqlStatement := `DELETE FROM product WHERE product_id=$1 and selled_at is null`
	if _, err := db.DB.Exec(sqlStatement, param.Id); err != nil {
		return err
	}
	return nil
}
