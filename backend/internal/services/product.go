package services

import (
	"secondChance/internal/models"
	"time"
)

func (s *Layer) GetAllProduct(b *models.IsAuction) ([]models.Products, error) {
	products, err := s.DBLayer.GetAllProduct(b)
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

func (s *Layer) GetProductAuction(id *models.IdReg) (*models.Product, error) {
	product, err := s.DBLayer.GetProductAuction(id)
	if err != nil {
		return &models.Product{}, err
	}
	return product, nil
}

func (s *Layer) SoldProduct(t time.Time, id int) error {
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

func (s *Layer) DeleteProduct(id *models.IdReg) error {
	if err := s.DBLayer.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

func (s *Layer) UpdateProduct(id *models.IdReg, productReq *models.Product) error {
	productDB, err := s.DBLayer.GetProduct(id)
	if err != nil {
		return err
	}

	product := newProduct(productReq, productDB)
	if err := s.DBLayer.UpdateProduct(id, product); err != nil {
		return err
	}

	return nil
}

func newProduct(productReq, productDB *models.Product) *models.Product {
	if productReq.Price != 0 {
		productDB.Price = productReq.Price
	}
	if productReq.Name != "" {
		productDB.Name = productReq.Name
	}
	if productReq.Description != "" {
		productDB.Description = productReq.Description
	}
	if productReq.Discount != 0 {
		productDB.Discount = productReq.Discount
	}
	if productReq.IsAuction.Auction != productDB.IsAuction.Auction {
		productDB.IsAuction.Auction = productReq.IsAuction.Auction
	}
	return productDB
}
