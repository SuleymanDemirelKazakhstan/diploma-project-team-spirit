package services

import (
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

func NewOwnerService(repo db.Shop) *OwnerService {
	return &OwnerService{
		repo: repo,
	}
}

func (o *OwnerService) GetAll() ([]models.Product, error) {
	products, err := o.repo.GetAll()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (o *OwnerService) Get(id *models.IdReg) (*models.Product, *models.Owner, error) {
	product, shop, err := o.repo.Get(id)
	if err != nil {
		return &models.Product{}, &models.Owner{}, err
	}
	return product, shop, nil
}

func (o *OwnerService) Create(product *models.CreateProduct) (*models.ImagePath, error) {
	path, err := o.repo.Create(product)
	if err != nil {
		return &models.ImagePath{}, err
	}
	return path, nil
}

func (o *OwnerService) Delete(id *models.IdReg) error {
	if err := o.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) Update(product *models.CreateProduct) (*models.ImagePath, error) {
	path, err := o.repo.Update(product)
	if err != nil {
		return &models.ImagePath{}, err
	}

	return path, nil
}

func (o *OwnerService) GetOrder(id *models.IdReg) (*[]models.OwnerOrder, error) {
	products, err := o.repo.GetOrder(id)
	if err != nil {
		return &[]models.OwnerOrder{}, err
	}
	return products, nil
}

func (o *OwnerService) Login(param *models.Login) (string, int, error) {
	owner, err := o.repo.GetOwner(param)
	if err != nil {
		return "", -1, err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = owner.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", -1, err
	}
	return t, owner.Id, nil
}

func (o *OwnerService) SaveImage(id *models.IdReg, file string) (string, error) {
	path, err := o.repo.SaveImage(id, file)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (o *OwnerService) DeleteImage(id *models.Image) error {
	if err := o.repo.DeleteImage(id); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) Issued(param *models.Issued) error {
	if err := o.repo.Issued(param); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) GetOrders(param *models.OwnerFillter) ([]models.OwnerProduct, error) {
	products, err := o.repo.GetOrders(param)
	if err != nil {
		return []models.OwnerProduct{}, err
	}
	return products, nil
}

func (o *OwnerService) GetCatalog(param *models.CatalogFilter) ([]models.Product, error) {
	products, err := o.repo.GetCatalog(param)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (o *OwnerService) UpdateEmail(param *models.EmailUser) error {
	if err := o.repo.UpdateEmail(param); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) GetProfile(param *models.IdReg) (*models.DTOowner, error) {
	user, err := o.repo.GetProfile(param)
	if err != nil {
		return &models.DTOowner{}, err
	}
	return user, nil
}

func (o *OwnerService) UpdatePassword(param *models.Password) error {
	if err := o.repo.UpdatePassword(param); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) UpdateProfile(param *models.DTOowner) error {
	if err := o.repo.UpdateProfile(param); err != nil {
		return err
	}
	return nil
}

func (o *OwnerService) MainPage(id *models.IdReg) (*models.MainPage, []models.OwnerProduct, error) {
	param, products, err := o.repo.MainPage(id)
	if err != nil {
		return &models.MainPage{}, []models.OwnerProduct{}, err
	}
	return param, products, nil
}
