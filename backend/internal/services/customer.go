package services

import "secondChance/internal/models"

func (s *Layer) GetCustomer(param string) (*models.Customer, error) {
	user, err := s.DBLayer.GetCustomer(param)
	if err != nil {
		return &models.Customer{}, err
	}
	return user, nil
}

func (s *Layer) CreateCustomer(user *models.Customer) error {
	if err := s.DBLayer.CreateCustomer(user); err != nil {
		return err
	}
	return nil
}
