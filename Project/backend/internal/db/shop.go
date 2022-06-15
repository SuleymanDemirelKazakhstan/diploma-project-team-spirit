package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"secondChance/internal/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type OwnerRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewOwnerRepo(db *sql.DB, rdb *redis.Client) *OwnerRepo {
	return &OwnerRepo{
		db:  db,
		rdb: rdb,
	}
}

func (o *OwnerRepo) Create(product *models.CreateProduct) (*models.ImagePath, error) {
	var id models.IdReg
	paths := new(models.ImagePath)
	sqlStatement := `select max(product_id) from product`

	row := o.db.QueryRow(sqlStatement)
	if err := row.Scan(&id.Id); err != nil {
		return &models.ImagePath{}, err
	}
	for _, v := range product.FileName {
		path := fmt.Sprintf("/images/product/%d/%s", id.Id+1, v)
		if err := os.MkdirAll(fmt.Sprintf("./images/product/%d", id.Id+1), os.ModePerm); err != nil {
			return &models.ImagePath{}, err
		}
		paths.Path = append(paths.Path, path)
	}

	tx, err := o.db.Begin()
	if err != nil {
		return &models.ImagePath{}, err
	}

	sqlStatement = `INSERT INTO product (shop_id, price, name, 
		description, is_auction, product_category, product_subcategory, 
		product_size, product_colour, discount, product_condition, image) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	if _, err := tx.Exec(sqlStatement, product.Id,
		product.Price, product.Name,
		product.Description, product.Auction,
		product.Category, product.Subcategory,
		product.Size, product.Colour, product.Discount,
		product.Condition, pq.Array(paths.Path)); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return &models.ImagePath{}, txErr
		}
		return &models.ImagePath{}, err
	}

	if product.Auction {
		loc, err := time.LoadLocation("Asia/Almaty")
		if err != nil {
			return &models.ImagePath{}, err
		}

		RFC3339local := "2006-01-02T15:04:05Z"
		t1, err := time.ParseInLocation(RFC3339local, product.Selled, loc)
		if err != nil {
			return &models.ImagePath{}, err
		}

		val, err := json.Marshal(models.Value{
			Price:      int(product.Price),
			CustomerId: -1,
			StartTime:  t1,
		})
		if err != nil {
			return &models.ImagePath{}, err
		}
		s2 := strconv.Itoa(id.Id + 1)

		sqlStatement = `UPDATE product SET end_date=$2 WHERE product_id=$1`
		if _, err := tx.Exec(sqlStatement, id.Id+1, t1); err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return &models.ImagePath{}, txErr
			}
			return &models.ImagePath{}, err
		}

		if err := o.rdb.Set(context.Background(), s2, val, 0).Err(); err != nil {
			return &models.ImagePath{}, err
		}
	}

	if txErr := tx.Commit(); txErr != nil {
		return &models.ImagePath{}, err
	}

	return paths, nil
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

func (o *OwnerRepo) Update(product *models.CreateProduct) (*models.ImagePath, error) {
	str := []string{}
	paths := new(models.ImagePath)
	if product.Price != 0 {
		str = append(str, fmt.Sprintf("price=%f", product.Price))
	}
	if product.Name != "" {
		str = append(str, fmt.Sprintf("name='%s'", product.Name))
	}
	if product.Description != "" {
		str = append(str, fmt.Sprintf("description='%s'", product.Description))
	}
	if product.Discount != 0 {
		str = append(str, fmt.Sprintf("discount=%d", product.Discount))
	}
	if product.Category != "" {
		str = append(str, fmt.Sprintf("product_category='%s'", product.Category))
	}
	if product.Subcategory != "" {
		str = append(str, fmt.Sprintf("product_subcategory='%s'", product.Subcategory))
	}
	if product.Size != "" {
		str = append(str, fmt.Sprintf("product_size='%s'", product.Size))
	}
	if product.Colour != "" {
		str = append(str, fmt.Sprintf("product_colour='%s'", product.Colour))
	}
	if product.Condition != "" {
		str = append(str, fmt.Sprintf("product_condition='%s'", product.Condition))
	}
	str = append(str, fmt.Sprintf("is_auction=%v", product.Auction))

	sqlStatement := ""
	if product.FileName == nil {
		sqlStatement = fmt.Sprintf(`UPDATE product SET %v WHERE product_id=$1`, strings.Join(str, ","))
		_, err := o.db.Exec(sqlStatement, product.Id)
		if err != nil {
			return &models.ImagePath{}, err
		}
	}

	if product.FileName != nil {
		for _, v := range product.FileName {
			path := fmt.Sprintf("/images/product/%d/%s", product.Id, v)
			paths.Path = append(paths.Path, path)
		}
		sqlStatement = `select image from product where product_id=$1`
		row := o.db.QueryRow(sqlStatement, product.Id)
		if err := row.Scan(pq.Array(&paths.OldPath)); err != nil {
			return &models.ImagePath{}, err
		}

		m := make(map[string]bool)
		for _, value := range paths.OldPath {
			m[value] = true
		}
		for _, value := range paths.Path {
			if m[value] {
				m[value] = false
			}
		}
		paths.OldPath = paths.OldPath[:0]

		for key, value := range m {
			if value {
				paths.OldPath = append(paths.OldPath, key)
			}
		}

		sqlStatement = fmt.Sprintf(`UPDATE product SET %v, image=$2 WHERE product_id=$1`, strings.Join(str, ","))
		_, err := o.db.Exec(sqlStatement, product.Id, pq.Array(paths.Path))
		if err != nil {
			return &models.ImagePath{}, fmt.Errorf("image proglem %w", err)
		}
	}

	return paths, nil
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
	sqlStatement := `select t1.name, t1.email, t2.product_id, t2.name, t2.product_size, t2.price, t2.image, t2.crated_at, t2.selled_at, t3.status 
	from customer t1, product t2, orders t3 
	where t1.customer_id = t3.customer_id and t2.product_id = t3.product_id and t3.order_id = $1;`

	rows, err := o.db.Query(sqlStatement, id.Id)
	if err != nil {
		return &[]models.OwnerOrder{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.OwnerOrder
		if err := rows.Scan(&product.CustomerName, &product.CustomerEmail, &product.ProductId, &product.ProductName,
			&product.Size, &product.Price, pq.Array(&product.Image), &product.Create_at, &product.Selled_at, &product.Status); err != nil {
			return &[]models.OwnerOrder{}, err
		}
		products = append(products, product)
	}
	return &products, nil
}

func (o *OwnerRepo) Issued(param *models.Issued) error {
	sqlStatement := `UPDATE orders SET status=$2 WHERE product_id=$1`
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
	path := fmt.Sprintf("/images/product/%d/%s", id.Id, file)
	if err := os.MkdirAll(fmt.Sprintf("./images/product/%d", id.Id), os.ModePerm); err != nil {
		return "", err
	}
	sqlStatement := fmt.Sprintf(`UPDATE product SET image=array_append(image, '%s') WHERE product_id=$1`, path)
	_, err := o.db.Exec(sqlStatement, id.Id)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (o *OwnerRepo) DeleteImage(param *models.Image) error {
	path := fmt.Sprintf("/images/product/%d/%s", param.Id, param.Name)
	sqlStatement := fmt.Sprintf(`UPDATE product SET image=array_remove(image, '%s') WHERE product_id=$1`, path)
	_, err := o.db.Exec(sqlStatement, param.Id)
	if err != nil {
		return err
	}
	return nil
}

func (o *OwnerRepo) GetOrders(param *models.OwnerFillter) ([]models.OwnerProduct, error) {
	var products []models.OwnerProduct
	sqlStatement := `SELECT t1.product_id, t1.price, t1.name, t1.selled_at, t1.is_auction, t3.name, t2.status, t2.order_id
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
			&product.Auction, &product.Customer, &product.Status, &product.OrderId); err != nil {
			return []models.OwnerProduct{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (o *OwnerRepo) GetCatalog(param *models.CatalogFilter) ([]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, image, product_category, product_subcategory, product_size, product_colour, is_auction from product where selled_at is null and shop_id=$1`

	if param.MinPrice != 0 || param.MaxPrice != 0 {
		if param.MinPrice != 0 && param.MaxPrice != 0 {
			sqlStatement += fmt.Sprintf(" and price BETWEEN %d AND %d", param.MinPrice, param.MaxPrice)
		} else if param.MinPrice != 0 {
			sqlStatement += fmt.Sprintf(" and price >= %d", param.MinPrice)
		} else {
			sqlStatement += fmt.Sprintf(" and price <= %d", param.MaxPrice)
		}
	}
	if param.Category != nil && param.Category[0] != "" {
		for i := range param.Category {
			param.Category[i] = fmt.Sprintf("'%s'", param.Category[i])
		}
		sqlStatement += fmt.Sprintf(" and product_category in (%s)", strings.Join(param.Category, ","))
	}
	if param.Subcategory != nil && param.Subcategory[0] != "" {
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
	}
	if !param.Auction {
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
			&product.Category, &product.Subcategory, &product.Size, &product.Colour, &product.Auction); err != nil {
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

func (o *OwnerRepo) GetProfile(param *models.IdReg) (*models.DTOowner, error) {
	var user models.DTOowner
	sqlStatement := `SELECT name,email,phone,address,image,social_network FROM shop WHERE shop_id=$1 and is_deleted=false`

	row := o.db.QueryRow(sqlStatement, param.Id)
	// unmarshal the row object to user
	var social sql.NullString
	if err := row.Scan(&user.Name, &user.Email, &user.Phone, &user.Address, &user.Image, &social); err != nil {
		return &models.DTOowner{}, err
	}
	user.Social = social.String
	if err := godotenv.Load(); err != nil {
		return &models.DTOowner{}, err
	}
	_url := os.Getenv("baseUrl")
	user.Image = _url + user.Image
	return &user, nil
}

func (o *OwnerRepo) UpdatePassword(param *models.Password) error {
	var password models.Password
	sqlStatement := `SELECT password FROM shop WHERE shop_id=$1`
	row := o.db.QueryRow(sqlStatement, param.Id)
	if err := row.Scan(&password.Old); err != nil {
		return err
	}

	if !CheckPasswordHash(param.Old, password.Old) {
		return fmt.Errorf("password hash error")
	}

	sqlStatement = `UPDATE shop SET password=$2 WHERE shop_id=$1`
	_, err := o.db.Exec(sqlStatement, param.Id, param.New)
	if err != nil {
		return err
	}

	return nil
}

func (o *OwnerRepo) UpdateProfile(param *models.DTOowner) error {
	str := []string{}
	if param.Name != "" {
		str = append(str, fmt.Sprintf("name='%s'", param.Name))
	}
	if param.Address != "" {
		str = append(str, fmt.Sprintf("address='%s'", param.Address))
	}
	if param.Phone != "" {
		str = append(str, fmt.Sprintf("phone='%s'", param.Phone))
	}
	if param.Social != "" {
		str = append(str, fmt.Sprintf("social_network='%s'", param.Social))
	}

	if len(str) == 0 {
		return fmt.Errorf("params empty")
	}
	sqlStatement := fmt.Sprintf(`UPDATE shop SET %v WHERE shop_id=$1`, strings.Join(str, ","))
	_, err := o.db.Exec(sqlStatement, param.Id)
	if err != nil {
		return err
	}

	return nil
}

func (o *OwnerRepo) MainPage(id *models.IdReg) (*models.MainPage, []models.OwnerProduct, error) {
	var products []models.OwnerProduct
	sqlStatement := `SELECT t1.product_id, t1.price, t1.name, t1.selled_at, t1.is_auction, t3.name, t2.status, t2.order_id
	from product t1, orders t2, customer t3 where t1.selled_at is not null and t1.shop_id=$1 and t1.product_id = t2.product_id and t2.customer_id = t3.customer_id ORDER BY t1.selled_at ASC limit 4`
	rows, err := o.db.Query(sqlStatement, id.Id)
	if err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.OwnerProduct
		if err := rows.Scan(&product.Id, &product.Price, &product.Name, &product.Selled_at,
			&product.Auction, &product.Customer, &product.Status, &product.OrderId); err != nil {
			return &models.MainPage{}, []models.OwnerProduct{}, err
		}
		products = append(products, product)
	}

	var param models.MainPage
	sqlStatement = `SELECT DISTINCT count(product_id) FROM orders where shop_id=$1;`
	row := o.db.QueryRow(sqlStatement, id.Id)
	if err := row.Scan(&param.Customers); err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, err
	}

	sqlStatement = `SELECT count(order_id) FROM orders where shop_id=$1;`
	row = o.db.QueryRow(sqlStatement, id.Id)
	if err := row.Scan(&param.Orders); err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, err
	}

	sqlStatement = `SELECT sum(price-discount) FROM product where shop_id=$1;`
	row = o.db.QueryRow(sqlStatement, id.Id)
	if err := row.Scan(&param.Earnings); err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, fmt.Errorf("database Earnings %w", err)
	}

	sqlStatement = `SELECT name FROM shop where shop_id=$1;`
	row = o.db.QueryRow(sqlStatement, id.Id)
	if err := row.Scan(&param.Name); err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, err
	}

	sqlStatement = `SELECT count(*) FROM product where shop_id=$1 and selled_at is null;`
	row = o.db.QueryRow(sqlStatement, id.Id)
	if err := row.Scan(&param.Products); err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, err
	}

	return &param, products, nil
}
