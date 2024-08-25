package role

import (
	"github.com/mrrizkin/finteligo/app/models"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(role *models.Role) (*models.Role, error) {
	err := s.repo.Create(role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *Service) FindAll() ([]models.Role, error) {
	roles, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *Service) FindByID(id uint) (*models.Role, error) {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *Service) Update(id uint, role *models.Role) (*models.Role, error) {
	var err error

	_, err = s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
