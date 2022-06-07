package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"secondChance/internal/models"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type CustomerRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewCustomerRepo(db *sql.DB, rdb *redis.Client) *CustomerRepo {
	return &CustomerRepo{
		db:  db,
		rdb: rdb,
	}
}

func (c *CustomerRepo) Get(email string) (*models.Customer, error) {
	var user models.Customer
	sqlStatement := `SELECT customer_id,name,email,password,phone, image FROM customer WHERE email=$1`

	row := c.db.QueryRow(sqlStatement, email)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Image); err != nil {
		return &models.Customer{}, err
	}

	if err := godotenv.Load(); err != nil {
		return &models.Customer{}, err
	}
	_url := os.Getenv("baseUrl")

	user.Image = _url + user.Image
	return &user, nil
}

func (c *CustomerRepo) GetPassword(email string) (*models.Customer, error) {
	var user models.Customer
	sqlStatement := `SELECT customer_id,name,email,password FROM customer WHERE email=$1`

	row := c.db.QueryRow(sqlStatement, email)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
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

	sqlStatement := `INSERT INTO orders (customer_id, product_id, shop_id) VALUES ($1, $2, $3)`
	if _, err := tx.Exec(sqlStatement, order.Customer_id, order.Product_id, order.Shop_id); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	sqlStatement = `UPDATE product SET selled_at=$2 WHERE product_id=$1`
	if _, err := tx.Exec(sqlStatement, order.Product_id, time.Now()); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}
	if txErr := tx.Commit(); txErr != nil {
		return txErr
	}
	return nil
}

func (c *CustomerRepo) GetOrder(id *models.IdReg) (*[]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, price, name, discount, selled_at from "product" where product_id in (select product_id from "orders" where customer_id = $1)`

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

func (c *CustomerRepo) SaveImage(id *models.IdReg, file string) (string, error) {
	// generate new uuid for image name
	uniqueId := uuid.New()
	// remove "- from imageName"
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	// extract image extension from original file filename
	fileExt := strings.Split(file, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	path := fmt.Sprintf("/images/customer/%d/%s", id.Id, image)
	if err := os.MkdirAll(fmt.Sprintf("./images/customer/%d", id.Id), os.ModePerm); err != nil {
		return "", err
	}

	sqlStatement := `UPDATE customer SET image=$2 WHERE customer_id=$1`
	_, err := c.db.Exec(sqlStatement, id.Id, path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *CustomerRepo) DeleteImage(id *models.IdReg) error {
	sqlStatement := `UPDATE customer SET image=NULL WHERE customer_id=$1`
	_, err := c.db.Exec(sqlStatement, id.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepo) Setter(deal *models.Deal, t time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	json, err := json.Marshal(models.Value{
		Price:      deal.Price,
		CustomerId: deal.CustomerId,
		StartTime:  time.Now(),
	})
	if err != nil {
		return err
	}

	if err := c.rdb.Set(ctx, deal.ProductId.Id, json, t).Err(); err != nil {
		return err
	}

	return nil
}

func (c *CustomerRepo) Getter(id *models.ProductId) (*models.Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	v, err := c.rdb.Get(ctx, id.Id).Result()
	if err != nil {
		return &models.Value{}, err
	}

	data := new(models.Value)
	if err := json.Unmarshal([]byte(v), data); err != nil {
		return &models.Value{}, err
	}
	return data, nil
}

func (c *CustomerRepo) GetFilter(f *models.Filter) ([]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, shop_id, price, name, image, discount, product_size, is_auction from product where selled_at is null`
	if f.Category != "" {
		sqlStatement += fmt.Sprintf(" and product_category = '%s'", f.Category)
	}
	if f.Subcategory != "" {
		sqlStatement += fmt.Sprintf(" and product_subcategory = '%s'", f.Subcategory)
	}
	if f.Size != "" {
		sqlStatement += fmt.Sprintf(" and product_size = '%s'", f.Size)
	}
	if f.Colour != "" {
		sqlStatement += fmt.Sprintf(" and product_colour = '%s'", f.Colour)
	}
	if f.Condition != "" {
		sqlStatement += fmt.Sprintf(" and product_condition = '%s'", f.Condition)
	}
	if f.MinPrice != 0 || f.MaxPrice != 0 {
		if f.MinPrice != 0 && f.MaxPrice != 0 {
			sqlStatement += fmt.Sprintf(" and price BETWEEN %d AND %d", f.MinPrice, f.MaxPrice)
		} else if f.MinPrice != 0 {
			sqlStatement += fmt.Sprintf(" and price >= %d", f.MinPrice)
		} else {
			sqlStatement += fmt.Sprintf(" and price <= %d", f.MaxPrice)
		}
	}
	fmt.Println(f.Type, reflect.TypeOf(f.Type))
	if f.Type > 0 {
		if f.Type == 1 {
			sqlStatement += " and is_auction=true"
		} else {
			sqlStatement += " and is_auction=false"
		}
	}
	fmt.Println(sqlStatement)
	rows, err := c.db.Query(sqlStatement)
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
		if err := rows.Scan(&product.Id, &product.OwnerId, &product.Price, &product.Name, pq.Array(&product.Image), &product.Discount, &product.Size, &product.Auction); err != nil {
			return []models.Product{}, err
		}
		for i := range product.Image {
			product.Image[i] = _url + product.Image[i]
		}
		products = append(products, product)
	}
	return products, nil
}

func (c *CustomerRepo) GetDiscountProducts() ([]models.Product, error) {
	var products []models.Product
	sqlStatement := `SELECT product_id, shop_id, price, name, image, discount, product_category, product_size from product where selled_at is null and discount != 0 limit 4`

	rows, err := c.db.Query(sqlStatement)
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
		if err := rows.Scan(&product.Id, &product.OwnerId, &product.Price,
			&product.Name, pq.Array(&product.Image), &product.Discount,
			&product.Category, &product.Size); err != nil {
			return []models.Product{}, err
		}
		for i := range product.Image {
			product.Image[i] = _url + product.Image[i]
		}
		products = append(products, product)
	}
	return products, nil
}

func (c *CustomerRepo) Search(p *models.SearchParam) ([]models.Product, error) {
	var products []models.Product
	sqlStatement := fmt.Sprintf(`SELECT product_id, shop_id, price, name, image, discount, product_category, product_size from product where selled_at is null and name like '%s%%';`, p.Param)

	rows, err := c.db.Query(sqlStatement)
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
		if err := rows.Scan(&product.Id, &product.OwnerId, &product.Price,
			&product.Name, pq.Array(&product.Image), &product.Discount,
			&product.Category, &product.Size); err != nil {
			return []models.Product{}, err
		}
		for i := range product.Image {
			product.Image[i] = _url + product.Image[i]
		}
		products = append(products, product)
	}
	return products, nil
}

func (c *CustomerRepo) GetAllMyProduct(id *models.IdReg) ([]models.CustomerOrder, error) {
	var products []models.CustomerOrder
	sqlStatement := `select t1.product_id, t1.selled_at, t2.address, t3.status from product t1, shop t2, orders t3 where t1.product_id = t3.product_id and t2.shop_id = t3.shop_id and t3.customer_id = $1;`

	rows, err := c.db.Query(sqlStatement, id.Id)
	if err != nil {
		return []models.CustomerOrder{}, err
	}
	defer rows.Close()

	if err := godotenv.Load(); err != nil {
		return []models.CustomerOrder{}, err
	}
	_url := os.Getenv("baseUrl")

	for rows.Next() {
		var product models.CustomerOrder
		if err := rows.Scan(&product.ProductId, &product.SelledAt,
			&product.Address, pq.Array(&product.Image),
			&product.Status); err != nil {
			return []models.CustomerOrder{}, err
		}
		for i := range product.Image {
			product.Image[i] = _url + product.Image[i]
		}
		products = append(products, product)
	}
	return products, nil
}
