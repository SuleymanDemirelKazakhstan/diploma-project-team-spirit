package services

import "secondChance/internal/models"

func (s *Layer) CreateOwner(user *models.Owner) error {
	if err := s.DBLayer.CreateOwner(user); err != nil {
		return err
	}
	return nil
}

func (s *Layer) DeleteOwner(param *models.OwnerEmailRequest) error {
	if err := s.DBLayer.DeleteOwner(param); err != nil {
		return err
	}
	return nil
}

func (s *Layer) GetOwner(param *models.OwnerEmailRequest) (*models.GetOwner, error) {
	user, err := s.DBLayer.GetOwner(param)
	if err != nil {
		return &models.GetOwner{}, err
	}
	return user, nil
}

func (s *Layer) GetAllOwner() ([]models.Owner, error) {
	users, err := s.DBLayer.GetAllOwner()
	if err != nil {
		return []models.Owner{}, err
	}
	return users, nil
}

func (s *Layer) UpdateOwner(email *models.OwnerEmailRequest, user *models.Owner) error {
	if err := s.DBLayer.UpdateOwner(email, user); err != nil {
		return err
	}
	return nil
}
