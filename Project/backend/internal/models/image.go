package models

type Image struct {
	Name string `json:"name,omitempty" validate:"required"`
	Id   int    `json:"id,omitempty" validate:"required"`
}
