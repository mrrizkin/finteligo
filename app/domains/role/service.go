package role

import (
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Create(payload *RolePayload) (*models.Role, error) {
	role := &models.Role{
		Name: payload.Name,
		Slug: payload.Slug,
	}

	err := s.repo.Create(role, payload.PermissionIDs)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *Service) FindAll(pagination types.Pagination) (*PaginatedRole, error) {
	roles, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	rolesCount, err := s.repo.FindAllCount()
	if err != nil {
		return nil, err
	}

	return &PaginatedRole{
		Result: roles,
		Total:  int(rolesCount),
	}, nil
}

func (s *Service) FindByID(id uint) (*models.Role, error) {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *Service) Update(id uint, role *RolePayload) (*models.Role, error) {
	exRole, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	exRole.Name = role.Name
	exRole.Slug = role.Slug

	err = s.repo.Update(exRole, role.PermissionIDs)
	if err != nil {
		return nil, err
	}

	return exRole, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
