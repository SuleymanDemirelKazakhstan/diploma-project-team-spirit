package db

import (
	"database/sql"
	"fmt"
	"os"
	"secondChance/internal/models"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type OwnerRepo struct {
	db *sql.DB
}

func NewOwnerRepo(db *sql.DB) *OwnerRepo {
	return &OwnerRepo{db: db}
}

func (o *OwnerRepo) Create(product *models.Product) error {
	sqlStatement := `INSERT INTO product (shop_id, price, name, description, is_auction, product_category, product_subcategory, product_size, product_colour, discount, product_condition) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if err := o.db.QueryRow(sqlStatement, product.OwnerId,
		product.Price, product.Name,
		product.Description, product.Auction,
		product.Category, product.Subcategory,
		product.Size, product.Colour, product.Discount,
		product.Condition); err != nil {
		return err.Err()
	}
	return nil
}

func (o *OwnerRepo) Get(id *models.IdReg) (*models.Product, *models.Owner, error) {
	var product models.Product
	var user models.Owner
	sqlStatement := `SELECT shop_id,price, name,description, discount, image, is_auction, product_category, product_subcategory, product_size, product_colour, product_condition FROM product WHERE product_id=$1`

	row := o.db.QueryRow(sqlStatement, id.Id)
	// unmarshal the row object to user
	if err := row.Scan(&product.OwnerId, &product.Price, &product.Name,
		&product.Description, &product.Discount,
		pq.Array(&product.Image), &product.Auction,
		&product.Category, &product.Subcategory,
		&product.Size, &product.Colour,
		&product.Condition); err != nil {
		return &models.Product{}, &models.Owner{}, err
	}
	if err := godotenv.Load(); err != nil {
		return &models.Product{}, &models.Owner{}, err
	}
	_url := os.Getenv("baseUrl")
	for i := range product.Image {
		product.Image[i] = _url + product.Image[i]
	}

	sqlStatement = `SELECT shop_id,name, description,email,phone,address FROM shop WHERE shop_id=$1 and is_deleted=false`
	row = o.db.QueryRow(sqlStatement, product.OwnerId)
	if err := row.Scan(&user.Id, &user.Name, &user.Description, &user.Email, &user.Phone, &user.Address); err != nil {
		return &models.Product{}, &models.Owner{}, err
	}

	return &product, &user, nil
}

func (o *OwnerRepo) GetAll() ([]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, shop_id, price, name, image from product where selled_at is null`

	rows, err := o.db.Query(sqlStatement)
	if err != nil {
		return []models.Product{}, err
	}
	defer rows.Close()
	if err := godotenv.Load(); err != nil {
		return []models.Product{}, err
	}
	_url := os.Getenv("baseUrl")
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.OwnerId, &product.Price, &product.Name, pq.Array(&product.Image)); err != nil {
			return []models.Product{}, err
		}
		for i := range product.Image {
			product.Image[i] = _url + product.Image[i]
		}
		products = append(products, product)
	}
	return products, nil
}

func (o *OwnerRepo) Update(product *models.Product) error {
	str := []string{}
	if product.Price != 0 {
		str = append(str, fmt.Sprintf("price=%v", product.Price))
	}
	if product.Name != "" {
		str = append(str, fmt.Sprintf("name=%v", product.Name))
	}
	if product.Description != "" {
		str = append(str, fmt.Sprintf("description=%v", product.Description))
	}
	if product.Discount != 0 {
		str = append(str, fmt.Sprintf("discount=%v", product.Discount))
	}
	str = append(str, fmt.Sprintf("is_auction=%v", product.Auction))

	sqlStatement := fmt.Sprintf(`UPDATE product SET %v WHERE product_id=$1`, strings.Join(str, ","))
	_, err := o.db.Exec(sqlStatement, product.Id, product.Price, product.Name,
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

func (o *OwnerRepo) GetOrder(id *models.IdReg) (*[]models.OwnerOrder, error) {
	var products []models.OwnerOrder
	sqlStatement := `select t1.name, t2.name, t2.price, t2.is_auction, t2.selled_at t3.status from customer t1, product t2, orders t3 where t1.customer_id = t3.customer_id and t2.product_id = t3.product_id and t3.shop_id = $1;`

	rows, err := o.db.Query(sqlStatement, id.Id)
	if err != nil {
		return &[]models.OwnerOrder{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.OwnerOrder
		if err := rows.Scan(&product.CustomerName, &product.ProductName, &product.Price, &product.Auction, &product.Selled_at, &product.Status); err != nil {
			return &[]models.OwnerOrder{}, err
		}
		products = append(products, product)
	}
	return &products, nil
}

func (o *OwnerRepo) Issued(param *models.Issued) error {
	sqlStatement := `UPDATE orders SET status=%2 WHERE product_id=$1`
	_, err := o.db.Exec(sqlStatement, param.Id, param.Issued)
	if err != nil {
		return err
	}
	return nil
}

func (o *OwnerRepo) GetOwner(param *models.Login) (*models.Owner, error) {
	var user models.Owner
	sqlStatement := `SELECT shop_id,email,password FROM shop WHERE email=$1`

	row := o.db.QueryRow(sqlStatement, param.Email)
	if err := row.Scan(&user.Id, &user.Email, &user.Password); err != nil {
		return &models.Owner{}, err
	}
	if !CheckPasswordHash(param.Password, user.Password) {
		return nil, fmt.Errorf("login error")
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
	path := fmt.Sprintf("/images/product/%d/%s", id.Id, image)
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

func (o *OwnerRepo) GetAllMyProduct(param *models.OwnerFillter) ([]models.OwnerProduct, error) {
	var products []models.OwnerProduct
	sqlStatement := `SELECT t1.product_id, t1.price, t1.name, t1.selled_at, t1.is_auction, t3.name, t2.status
	from product t1, orders t2, customer t3 where t1.selled_at is not null and t1.shop_id=$1 and t1.product_id = t2.product_id and t2.customer_id = t3.customer_id`

	if param.Status > 0 {
		if param.Status == 1 {
			sqlStatement += " and t2.status=true"
		} else {
			sqlStatement += " and t2.status=false"
		}
	}
	if param.MinPrice != 0 || param.MaxPrice != 0 {
		if param.MinPrice != 0 && param.MaxPrice != 0 {
			sqlStatement += fmt.Sprintf(" and t1.price BETWEEN %d AND %d", param.MinPrice, param.MaxPrice)
		} else if param.MinPrice != 0 {
			sqlStatement += fmt.Sprintf(" and t1.price >= %d", param.MinPrice)
		} else {
			sqlStatement += fmt.Sprintf(" and t1.price <= %d", param.MaxPrice)
		}
	}
	if !param.StartDate.IsZero() || !param.EndDate.IsZero() {
		if !param.StartDate.IsZero() && !param.EndDate.IsZero() {
			sqlStatement += fmt.Sprintf(" and t1.selled_at BETWEEN %d AND %d", param.StartDate, param.EndDate)
		} else if !param.StartDate.IsZero() {
			sqlStatement += fmt.Sprintf(" and t1.selled_at >= %d", param.StartDate)
		} else {
			sqlStatement += fmt.Sprintf(" and t1.selled_at <= %d", param.EndDate)
		}
	}
	if param.Search != "" {
		sqlStatement += fmt.Sprintf(" and t1.name like '%s%%'", param.Search)
	}

	rows, err := o.db.Query(sqlStatement, param.Id)
	if err != nil {
		return []models.OwnerProduct{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.OwnerProduct
		if err := rows.Scan(&product.Id, &product.Price, &product.Name, &product.Selled_at,
			&product.Auction, &product.Customer, &product.Status); err != nil {
			return []models.OwnerProduct{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (o *OwnerRepo) GetCatalog(param *models.CatalogFilter) ([]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, image, product_category, product_subcategory, product_size, product_colour from product where selled_at is null and shop_id=$1`

	if param.MinPrice != 0 || param.MaxPrice != 0 {
		if param.MinPrice != 0 && param.MaxPrice != 0 {
			sqlStatement += fmt.Sprintf(" and price BETWEEN %d AND %d", param.MinPrice, param.MaxPrice)
		} else if param.MinPrice != 0 {
			sqlStatement += fmt.Sprintf(" and price >= %d", param.MinPrice)
		} else {
			sqlStatement += fmt.Sprintf(" and price <= %d", param.MaxPrice)
		}
	}
	if param.Category != nil {
		for i := range param.Category {
			param.Category[i] = fmt.Sprintf("'%s'", param.Category[i])
		}
		sqlStatement += fmt.Sprintf(" and product_category in (%s)", strings.Join(param.Category, ","))
		fmt.Println(sqlStatement)
	}
	if param.Subcategory != nil {
		for i := range param.Subcategory {
			param.Subcategory[i] = fmt.Sprintf("'%s'", param.Subcategory[i])
		}
		sqlStatement += fmt.Sprintf(" and product_subcategory in (%s)", strings.Join(param.Subcategory, ","))
	}
	if param.Search != "" {
		sqlStatement += fmt.Sprintf(" and name like '%s%%'", param.Search)
	}
	if param.Auction {
		sqlStatement += " and is_auction=true"
	} else {
		sqlStatement += " and is_auction=false"
	}

	rows, err := o.db.Query(sqlStatement, param.Id)
	if err != nil {
		return []models.Product{}, err
	}
	defer rows.Close()

	if err := godotenv.Load(); err != nil {
		return []models.Product{}, err
	}
	_url := os.Getenv("baseUrl")

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Price, &product.Name, pq.Array(&product.Image),
			&product.Category, &product.Subcategory, &product.Size, &product.Colour); err != nil {
			return []models.Product{}, err
		}
		for i := range product.Image {
			product.Image[i] = _url + product.Image[i]
		}
		products = append(products, product)
	}
	return products, nil
}

func (o *OwnerRepo) UpdateEmail(param *models.EmailUser) error {
	sqlStatement := `UPDATE shop SET email=$2 WHERE shop_id=$1`
	_, err := o.db.Exec(sqlStatement, param.Id, param.Email)
	if err != nil {
		return err
	}

	return nil
}
