package models

type Customer struct {
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Phone    string `json:"phone,omitempty" validate:"required"`
	Image    string `json:"image"       form:"image"`
}

type SearchParam struct {
	Param string `json:"param,omitempty" validate:"required"`
}

type Password struct {
	Id  int    `json:"id,omitempty" validate:"required"`
	Old string `json:"old,omitempty" validate:"required"`
	New string `json:"new,omitempty" validate:"required"`
}
