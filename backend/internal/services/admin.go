package services

import (
	"secondChance/internal/models"
)

func (s *Layer) CreateOwner(user *models.Owner) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
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

func (s *Layer) GetOwner(param *models.OwnerEmailRequest) (*models.Owner, error) {
	user, err := s.DBLayer.GetOwner(param)
	if err != nil {
		return &models.Owner{}, err
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

func (s *Layer) UpdateOwner(email *models.OwnerEmailRequest, userReq *models.Owner) error {
	userDB, err := s.DBLayer.GetOwner(email)
	if err != nil {
		return err
	}
	user, err := newUser(userReq, userDB)

	if err := s.DBLayer.UpdateOwner(email, user); err != nil {
		return err
	}
	return nil
}

//Todo validate
func newUser(userReq, user *models.Owner) (*models.Owner, error) {
	if userReq.Email != "" {
		user.Email = userReq.Email
	}
	if userReq.Name != "" {
		user.Name = userReq.Name
	}
	if userReq.Phone != 0 {
		user.Phone = userReq.Phone
	}
	if userReq.Address != "" {
		user.Address = userReq.Address
	}
	if userReq.Password != "" {
		var err error
		user.Password, err = HashPassword(userReq.Password)
		if err != nil {
			return &models.Owner{}, err
		}
	}
	return user, nil
}
