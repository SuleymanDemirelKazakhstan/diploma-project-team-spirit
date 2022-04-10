package services

import (
	"secondChance/internal/models"
	"time"
)

func (s *Layer) GetAllProduct() ([]models.Products, error) {
	products, err := s.DBLayer.GetAllProduct()
	if err != nil {
		return []models.Products{}, err
	}
	return products, nil
}

func (s *Layer) GetProduct(id *models.IdReg) (*models.Product, error) {
	product, err := s.DBLayer.GetProduct(id)
	if err != nil {
		return &models.Product{}, err
	}
	return product, nil
}

func (s *Layer) SoldProduct(t time.Time, id models.IdReg) error {
	if err := s.DBLayer.SoldProduct(t, id); err != nil {
		return err
	}
	return nil
}

func (s *Layer) CreateProduct(product *models.Product) (err error) {
	if err := s.DBLayer.CreateProduct(product); err != nil {
		return err
	}
	return nil
}

func (s *Layer) UpdateProduct(id *models.IdReg, product *models.Product) error {
	if err := s.DBLayer.UpdateProduct(id, product); err != nil {
		return err
	}
	return nil
}
