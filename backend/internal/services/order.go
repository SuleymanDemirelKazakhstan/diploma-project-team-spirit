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

func (s *Layer) CustomerOrder(id *models.IdReg) (*[]models.Product, error) {
	products, err := s.DBLayer.CustomerOrder(id)
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}

func (s *Layer) OwnerOrder(id *models.IdReg) (*[]models.Product, error) {
	products, err := s.DBLayer.OwnerOrder(id)
	if err != nil {
		return &[]models.Product{}, err
	}
	return products, nil
}
