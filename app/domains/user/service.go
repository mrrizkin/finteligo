package user

import (
	"github.com/mrrizkin/finteligo/app/models"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(user *models.User) (*models.User, error) {
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindAll() ([]models.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) FindByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Update(id uint, user *models.User) (*models.User, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
