package services

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"secondChance/internal/models"
	"time"
)

func (s *Layer) LoginCustomer(param *models.LoginInput) (t string, id int, err error) {
	customer, err := s.DBLayer.GetCustomer(param.Email)
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

func (s *Layer) LoginOwner(param *models.LoginInput) (string, error) {
	owner, err := s.DBLayer.GetOwner(&models.OwnerEmailRequest{
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
