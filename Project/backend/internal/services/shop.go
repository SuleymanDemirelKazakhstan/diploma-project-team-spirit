package services

import (
	"context"
	"encoding/json"
	"os"
	"secondChance/internal/db"
	"secondChance/internal/models"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
)

type OwnerService struct {
	repo db.Shop
	rdb  *redis.Client
}

func NewOwnerService(repo db.Shop, rdb *redis.Client) *OwnerService {
	return &OwnerService{
		repo: repo,
		rdb:  rdb,
	}
}

func (o *OwnerService) GetAll() ([]models.Products, error) {
	products, err := o.repo.GetAll()
	if err != nil {
		return []models.Products{}, err
	}
	return products, nil
}

func (o *OwnerService) Get(id *models.IdReg) (*models.Product, error) {
	product, err := o.repo.Get(id)
	if err != nil {
		return &models.Product{}, err
	}
	return product, nil
}

func (o *OwnerService) Create(product *models.Product) error {
	//TODO: transaction
	if product.Auction {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		exp := time.Hour * 3
		json, err := json.Marshal(models.Value{
			Price:      int(product.Price),
			CustomerId: -1,
			StartTime:  time.Now(),
		})
		if err != nil {
			return err
		}

		if err := o.rdb.Set(ctx, string(product.Id), json, exp).Err(); err != nil {
			return err
		}
	}
	if err := o.repo.Create(product); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) Delete(id *models.IdReg) error {
	if err := o.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) Update(productReq *models.Product) error {
	id := &models.IdReg{productReq.Id}
	productDB, err := o.repo.Get(id)
	if err != nil {
		return err
	}

	product := newProduct(productReq, productDB)
	if err := o.repo.Update(id, product); err != nil {
		return err
	}

	return nil
}

func newProduct(productReq, productDB *models.Product) *models.Product {
	if productReq.Price != 0 {
		productDB.Price = productReq.Price
	}
	if productReq.Name != "" {
		productDB.Name = productReq.Name
	}
	if productReq.Description != "" {
		productDB.Description = productReq.Description
	}
	if productReq.Discount != 0 {
		productDB.Discount = productReq.Discount
	}
	if productReq.Auction != productDB.Auction {
		productDB.Auction = productReq.Auction
	}
	return productDB
}

func (o *OwnerService) GetOrder(id *models.IdReg) (*[]models.Product, error) {
	products, err := o.repo.GetOrder(id)
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}

func (o *OwnerService) Login(param *models.LoginInput) (string, error) {
	// s.DBLayer.GetOwner
	owner, err := o.repo.GetOwner(param.Email)
	if err != nil {
		return "", err
	}

	if !CheckPasswordHash(param.Password, owner.Password) {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = owner.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (o *OwnerService) SaveImage(id *models.IdReg, file string) (string, error) {
	path, err := o.repo.SaveImage(id, file)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (o *OwnerService) DeleteImage(id *models.IdReg) error {
	if err := o.repo.DeleteImage(id); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) Issued(id *models.IdReg) error {
	if err := o.repo.Issued(id); err != nil {
		return err
	}
	return nil
}
