package db

import (
	"database/sql"
	"secondChance/internal/models"
)

type Admin interface {
	Create(user *models.Owner) error
	Get(param *models.OwnerEmailRequest) (*models.Owner, error)
	GetAll() ([]models.Owner, error)
	Update(email *models.OwnerEmailRequest, user *models.Owner) error
	Delete(param *models.OwnerEmailRequest) error
}

type Customer interface {
	Get(email string) (*models.Customer, error)
	Create(user *models.Customer) error
	CreateOrder(order *models.Order) error
	GetOrder(id *models.IdReg) (*[]models.Product, error)
}

type Shop interface {
	Create(product *models.Product) error
	Get(id *models.IdReg) (*models.Product, error)
	GetAll() ([]models.Products, error)
	Update(id *models.IdReg, product *models.Product) error
	Delete(param *models.IdReg) error
	GetOrder(id *models.IdReg) (*[]models.Product, error)
	GetOwner(email string) (*models.Owner, error)
}

type Repository struct {
	Admin
	Customer
	Shop
}

func NewDataBaseLayers(db *sql.DB) *Repository {
	return &Repository{
		Admin:    NewAdminRepo(db),
		Customer: NewCustomerRepo(db),
		Shop:     NewOwnerRepo(db),
	}
}
