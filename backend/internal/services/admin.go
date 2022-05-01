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
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	if err := r.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (r *AdminService) Delete(param *models.OwnerEmailRequest) error {
	if err := r.repo.Delete(param); err != nil {
		return err
	}
	return nil
}

func (r *AdminService) Get(param *models.OwnerEmailRequest) (*models.Owner, error) {
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

func (r *AdminService) Update(email *models.OwnerEmailRequest, userReq *models.Owner) error {
	userDB, err := r.repo.Get(email)
	if err != nil {
		return err
	}
	user, err := newUser(userReq, userDB)

	if err := r.repo.Update(email, user); err != nil {
		return err
	}
	return nil
}

//Todo validate
func newUser(userReq, user *models.Owner) (*models.Owner, error) {
	if userReq.Email != "" {
		user.Email = userReq.Email
	}
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Phone != 0 {
		user.Phone = userReq.Phone
	}
	if userReq.Address != "" {
		user.Address = userReq.Address
	}
	if userReq.Password != "" {
		var err error
		user.Password, err = HashPassword(userReq.Password)
		if err != nil {
			return &models.Owner{}, err
		}
	}
	return user, nil
}

func (r *AdminService) Login(param *models.LoginInput) (string, error) {
	owner, err := r.repo.Get(&models.OwnerEmailRequest{
		Email: param.Email,
	})
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
