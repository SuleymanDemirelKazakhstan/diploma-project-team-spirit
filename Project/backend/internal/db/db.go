package db

import (
	"database/sql"
	"secondChance/internal/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type Admin interface {
	Create(user *models.Owner) error
	Get(param *models.IdReg) (*models.Owner, error)
	GetAll() ([]models.Owner, error)
	GetLogin(param *models.Login) (*models.Owner, error)
	Update(user *models.Owner) error
	Delete(param *models.IdReg) error
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
}

type Customer interface {
	Get(param *models.Login) (*models.Customer, error)
	Create(user *models.Customer) error
	CreateOrder(order *models.Order) error
	GetOrder(id *models.IdReg) (*[]models.Product, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
	Setter(deal *models.Deal, t time.Duration) error
	Getter(id *models.ProductId) (*models.Value, error)
	GetFilter(f *models.Filter) ([]models.Product, error)
	GetDiscountProducts() ([]models.Product, error)
	Search(p *models.SearchParam) ([]models.Product, error)
	GetAllMyProduct(id *models.IdReg) ([]models.CustomerOrder, error)
	UpdatePassword(param *models.Password) error
}

type Shop interface {
	Create(product *models.Product) error
	Get(id *models.IdReg) (*models.Product, *models.Owner, error)
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(param *models.IdReg) error
	GetOrder(id *models.IdReg) (*[]models.OwnerOrder, error)
	GetOwner(param *models.Login) (*models.Owner, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
	Issued(id *models.IdReg) error
	GetAllMyProduct(id *models.IdReg) ([]models.Product, error)
}

type Repository struct {
	Admin
	Customer
	Shop
}

func NewDataBaseLayers(db *sql.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Admin:    NewAdminRepo(db),
		Customer: NewCustomerRepo(db, rdb),
		Shop:     NewOwnerRepo(db),
	}
}
