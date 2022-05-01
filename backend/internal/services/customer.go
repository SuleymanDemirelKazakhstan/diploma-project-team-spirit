package services

import (
	"os"
	"secondChance/internal/db"
	"secondChance/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomerService struct {
	repo db.Customer
}

func NewCustomerService(repo db.Customer) *CustomerService {
	return &CustomerService{repo: repo}
}

func (c *CustomerService) Get(param string) (*models.Customer, error) {
	user, err := c.repo.Get(param)
	if err != nil {
		return &models.Customer{}, err
	}
	return user, nil
}

func (c *CustomerService) Create(user *models.Customer) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	if err := c.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) Login(param *models.LoginInput) (t string, id int, err error) {
	customer, err := c.repo.Get(param.Email)
	if err != nil {
		return "", 0, err
	}

	if !CheckPasswordHash(param.Password, customer.Password) {
		return "", 0, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = customer.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err = token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", 0, err
	}
	return t, customer.Id, nil
}

func (c *CustomerService) CreateOrder(order *models.Order) (err error) {
	if err := c.repo.CreateOrder(order); err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) GetOrder(id *models.IdReg) (*[]models.Product, error) {
	products, err := c.repo.GetOrder(id)
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}
