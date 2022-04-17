package handlers

import (
	"github.com/golang-jwt/jwt/v4"
)

func ValidToken(t *jwt.Token, name string) bool {
	claims := t.Claims.(jwt.MapClaims)
	if claims["name"] != name {
		return false
	}
	return true
}
