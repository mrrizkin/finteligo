package user

import (
	"errors"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/types"
	"github.com/mrrizkin/finteligo/third_party/argon2"
)

func NewService(repo *Repo, argon2 *argon2.Argon2) *Service {
	return &Service{repo, argon2}
}

func (s *Service) Create(user *models.User) (*models.User, error) {
	if user.Password == nil {
		return nil, errors.New("password is required")
	}

	hash, err := s.argon2.GenerateHashPassword(*user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = &hash
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindAll(pagination types.Pagination) (*PaginatedUser, error) {
	users, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	usersCount, err := s.repo.FindAllCount()
	if err != nil {
		return nil, err
	}

	return &PaginatedUser{
		Result: users,
		Total:  int(usersCount),
	}, nil
}

func (s *Service) FindByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Update(id uint, user *models.User) (*models.User, error) {
	userExist, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user.Password != nil {
		if *user.Password != "" {
			hash, err := s.argon2.GenerateHashPassword(*user.Password)
			if err != nil {
				return nil, err
			}

			userExist.Password = &hash
		}
	}

	userExist.Name = user.Name
	userExist.Email = user.Email
	userExist.Username = user.Username
	userExist.RoleID = user.RoleID

	err = s.repo.Update(userExist)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
