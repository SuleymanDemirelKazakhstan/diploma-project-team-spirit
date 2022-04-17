package services

import (
	"secondChance/internal/models"
	"time"
)

func (s *Layer) CreateOrder(order *models.Order) (err error) {
	if err := s.DBLayer.CreateOrder(order); err != nil {
		return err
	}
	if err := s.DBLayer.SoldProduct(time.Now(), int(order.Product_id)); err != nil {
		return err
	}
	return nil
}
