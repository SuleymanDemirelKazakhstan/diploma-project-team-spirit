package models

type Owner struct {
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Phone    string `json:"phone,omitempty" validate:"required"`
	Address  string `json:"address,omitempty" validate:"required"`
	Image    string `json:"image"       form:"image"`
}