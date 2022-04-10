package models

type ErrorResp struct {
	Status  bool   `json:"status" example:"false"`
	Message string `json:"message" example:"error cause"`
}
