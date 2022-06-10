package services

import (
	"secondChance/internal/db"
	"secondChance/internal/models"

	"github.com/go-redis/redis/v8"
)

type Admin interface {
	Create(user *models.Owner) error
	Delete(param *models.IdReg) error
	Get(param *models.IdReg) (*models.Owner, error)
	GetAll() ([]models.Owner, error)
	Update(userReq *models.Owner) error
	Login(param *models.Login) (string, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
}

type Customer interface {
	Create(user *models.Customer) error
	Login(param *models.Login) (t string, id int, err error)
	CreateOrder(order *models.Order) (err error)
	GetOrder(id *models.IdReg) (*[]models.Product, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
	GmailCode(email *models.EmailRequest) (int, error)
	Setter(deal *models.Deal) error
	Getter(id *models.ProductId) (*models.Value, error)
	GetFilter(f *models.Filter) ([]models.Product, error)
	GetDiscountProducts() ([]models.Product, error)
	Search(p *models.SearchParam) ([]models.Product, error)
	GetAllMyProduct(id *models.IdReg) ([]models.CustomerOrder, error)
	UpdatePassword(param *models.Password) error
	UpdateEmail(param *models.EmailUser) error
}

type Shop interface {
	GetAll() ([]models.Product, error)
	Get(id *models.IdReg) (*models.Product, *models.Owner, error)
	Create(product *models.Product) error
	Delete(id *models.IdReg) error
	Update(productReq *models.Product) error
	GetOrder(id *models.IdReg) (*[]models.OwnerOrder, error)
	Login(param *models.Login) (string, int, error)
	SaveImage(id *models.IdReg, file string) (string, error)
	DeleteImage(id *models.IdReg) error
	Issued(id *models.IdReg) error
	GetAllMyProduct(param *models.OwnerFillter) ([]models.OwnerProduct, error)
}

type Service struct {
	Admin
	Customer
	Shop
}

func NewService(repo *db.Repository, rdb *redis.Client) *Service {
	return &Service{
		Admin:    NewAdminService(repo.Admin),
		Customer: NewCustomerService(repo.Customer),
		Shop:     NewOwnerService(repo.Shop, rdb),
	}
}
