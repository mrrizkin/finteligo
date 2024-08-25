package role_permission

import (
	"github.com/mrrizkin/finteligo/app/models"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(role_permission *models.RolePermission) (*models.RolePermission, error) {
	err := s.repo.Create(role_permission)
	if err != nil {
		return nil, err
	}

	return role_permission, nil
}

func (s *Service) FindAll() ([]models.RolePermission, error) {
	role_permissions, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return role_permissions, nil
}

func (s *Service) FindByID(id uint) (*models.RolePermission, error) {
	role_permission, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return role_permission, nil
}

func (s *Service) Update(
	id uint,
	role_permission *models.RolePermission,
) (*models.RolePermission, error) {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(role_permission)
	if err != nil {
		return nil, err
	}

	return role_permission, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
