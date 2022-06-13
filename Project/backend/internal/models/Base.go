package models

type Resp struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"error cause"`
}

type Password struct {
	Id  int    `json:"id,omitempty" validate:"required"`
	Old string `json:"old,omitempty" validate:"required"`
	New string `json:"new,omitempty" validate:"required"`
}
