package models

type Owner struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Phone    int    `json:"phone,omitempty" validate:"required"`
}

type GetOwner struct {
	Id int `json:"id"`
	Owner
}

type OwnerEmailRequest struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}
