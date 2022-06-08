package models

type Login struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type EmailRequest struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}
