package services

import "secondChance/internal/db"

// Layer - databases management
type Layer struct {
	DBLayer *db.Layer
}

// NewServiceLayer - init new services Layer
func NewServiceLayer(db *db.Layer) *Layer {
	return &Layer{
		DBLayer: db,
	}
}
