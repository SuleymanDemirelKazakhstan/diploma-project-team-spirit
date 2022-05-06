package db

import (
	"database/sql"
	"fmt"
	"os"
	"secondChance/internal/models"
	"strings"

	"github.com/google/uuid"
)

type OwnerRepo struct {
	db *sql.DB
}

func NewOwnerRepo(db *sql.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}

func (o *OwnerRepo) Create(product *models.Product) error {
	sqlStatement := `INSERT INTO product (shop_id, price, name, description, is_auction) VALUES ($1, $2, $3, $4, $5)`
	if err := o.db.QueryRow(sqlStatement, product.OwnerId, product.Price, product.Name, product.Description, product.Auction); err != nil {
		return err.Err()
	}
	return nil
}

func (o *OwnerRepo) Get(id *models.IdReg) (*models.Product, error) {
	var product models.Product
	sqlStatement := `SELECT shop_id,price, name,description, discount FROM product WHERE product_id=$1`

	row := o.db.QueryRow(sqlStatement, id.Id)
	// unmarshal the row object to user
	if err := row.Scan(&product.OwnerId, &product.Price, &product.Name, &product.Description, &product.Discount); err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}

func (o *OwnerRepo) GetAll() ([]models.Products, error) {
	var products []models.Products
	sqlStatement := `SELECT product_id, shop_id, price, name from product where selled_at is null`

	rows, err := o.db.Query(sqlStatement)
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

func (o *OwnerRepo) Update(id *models.IdReg, product *models.Product) error {
	sqlStatement := `UPDATE product SET price=$2, name=$3, description=$4, discount=$5, is_auction=$6 WHERE id=$1`
	_, err := o.db.Exec(sqlStatement, id.Id, product.Price, product.Name,
		product.Description, product.Discount, product.Auction)
	if err != nil {
		return err
	}
	return nil
}

func (o *OwnerRepo) Delete(param *models.IdReg) error {
	sqlStatement := `DELETE FROM product WHERE product_id=$1 and selled_at is null`
	if _, err := o.db.Exec(sqlStatement, param.Id); err != nil {
		return err
	}
	return nil
}

func (o *OwnerRepo) GetOrder(id *models.IdReg) (*[]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, discount, selled_at from "product" where product_id in (select product_id from "order" where shop_id = $1)`

	rows, err := o.db.Query(sqlStatement, id.Id)
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

func (o *OwnerRepo) GetOwner(email string) (*models.Owner, error) {
	var user models.Owner
	sqlStatement := `SELECT shop_id,email,password FROM shop WHERE email=$1`

	row := o.db.QueryRow(sqlStatement, email)
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return &models.Owner{}, err
	}
	return &user, nil
}

func (o *OwnerRepo) SaveImage(id *models.IdReg, file string) (string, error) {
	// generate new uuid for image name
	uniqueId := uuid.New()
	// remove "- from imageName"
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	// extract image extension from original file filename
	fileExt := strings.Split(file, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	path := fmt.Sprintf("./images/product/%d/%s", id.Id, image)
	if err := os.MkdirAll(fmt.Sprintf("./images/product/%d", id.Id), os.ModePerm); err != nil {
		return "", err
	}

	sqlStatement := `UPDATE product SET image=$2 WHERE product_id=$1`
	_, err := o.db.Exec(sqlStatement, id.Id, path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (o *OwnerRepo) DeleteImage(id *models.IdReg) error {
	sqlStatement := `UPDATE product SET image=NULL WHERE product_id=$1`
	_, err := o.db.Exec(sqlStatement, id.Id)
	if err != nil {
		return err
	}
	return nil
}
