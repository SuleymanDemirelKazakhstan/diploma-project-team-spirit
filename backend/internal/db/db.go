package db

import (
	"database/sql"
)

type Layer struct {
	DB *sql.DB
}

func NewDataBaseLayers(db *sql.DB) *Layer {
	return &Layer{
		DB: db,
	}
}
