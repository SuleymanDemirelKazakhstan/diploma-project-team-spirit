package services

import (
	"os"
	"secondChance/internal/db"
	"secondChance/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AdminService struct {
	repo db.Admin
}

func NewAdminService(repo db.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (r *AdminService) Create(user *models.Owner) error {
	if err := r.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (r *AdminService) Delete(param *models.IdReg) error {
	if err := r.repo.Delete(param); err != nil {
		return err
	}
	return nil
}

func (r *AdminService) Get(param *models.IdReg) (*models.Owner, error) {
	user, err := r.repo.Get(param)
	if err != nil {
		return &models.Owner{}, err
	}
	return user, nil
}

func (r *AdminService) GetAll() ([]models.Owner, error) {
	users, err := r.repo.GetAll()
	if err != nil {
		return []models.Owner{}, err
	}
	return users, nil
}

func (r *AdminService) Update(userReq *models.Owner) error {
	if err := r.repo.Update(userReq); err != nil {
		return err
	}
	return nil
}

func (r *AdminService) Login(param *models.Login) (string, error) {
	owner, err := r.repo.GetLogin(param)
	if err != nil {
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

func (r *AdminService) SaveImage(id *models.IdReg, file string) (string, error) {
	path, err := r.repo.SaveImage(id, file)
	if err != nil {
		return "", err
	}
	return path, nil
}
func (r *AdminService) DeleteImage(id *models.IdReg) error {
	if err := r.repo.DeleteImage(id); err != nil {
		return err
	}
	return nil
}
