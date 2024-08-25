package permission

import (
	"github.com/mrrizkin/finteligo/app/models"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(permission *models.Permission) (*models.Permission, error) {
	err := s.repo.Create(permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *Service) FindAll() ([]models.Permission, error) {
	permissions, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (s *Service) FindByID(id uint) (*models.Permission, error) {
	permission, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *Service) Update(id uint, permission *models.Permission) (*models.Permission, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
