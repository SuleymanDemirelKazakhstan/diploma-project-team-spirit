package services

import (
	"secondChance/internal/db"
	"secondChance/internal/models"
)

type Admin interface {
	Create(user *models.Owner) error
	Delete(param *models.OwnerEmailRequest) error
	Get(param *models.OwnerEmailRequest) (*models.Owner, error)
	GetAll() ([]models.Owner, error)
	Update(email *models.OwnerEmailRequest, userReq *models.Owner) error
	Login(param *models.LoginInput) (string, error)
}

type Customer interface {
	Get(param string) (*models.Customer, error)
	Create(user *models.Customer) error
	Login(param *models.LoginInput) (t string, id int, err error)
	CreateOrder(order *models.Order) (err error)
	GetOrder(id *models.IdReg) (*[]models.Product, error)
}

type Shop interface {
	GetAll() ([]models.Products, error)
	Get(id *models.IdReg) (*models.Product, error)
	Create(product *models.Product) error
	Delete(id *models.IdReg) error
	Update(id *models.IdReg, productReq *models.Product) error
	GetOrder(id *models.IdReg) (*[]models.Product, error)
	Login(param *models.LoginInput) (string, error)
}

type Service struct {
	Admin
	Customer
	Shop
}

func NewService(repo *db.Repository) *Service {
	return &Service{
		Admin:    NewAdminService(repo.Admin),
		Customer: NewCustomerService(repo.Customer),
		Shop:     NewOwnerService(repo.Shop),
	}
}
