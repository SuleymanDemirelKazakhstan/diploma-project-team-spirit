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
	GetLogin(param *models.EmailRequest) (*models.Owner, error)
	Update(user *models.Owner) error
	Delete(param *models.IdReg) error
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
}

type Customer interface {
	Get(email string) (*models.Customer, error)
	GetPassword(email string) (*models.Customer, error)
	Create(user *models.Customer) error
	CreateOrder(order *models.Order) error
	GetOrder(id *models.IdReg) (*[]models.Product, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
	Setter(deal *models.Deal, t time.Duration) error
	Getter(id *models.ProductId) (*models.Value, error)
	GetFilter(f *models.Filter) ([]models.Product, error)
}

type Shop interface {
	Create(product *models.Product) error
	Get(id *models.IdReg) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Update(id *models.IdReg, product *models.Product) error
	Delete(param *models.IdReg) error
	GetOrder(id *models.IdReg) (*[]models.Product, error)
	GetOwner(email string) (*models.Owner, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
	Issued(id *models.IdReg) error
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
