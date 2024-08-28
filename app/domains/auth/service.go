package auth

import (
	"errors"

	"github.com/mrrizkin/finteligo/app/models"
)

func NewService(repo *Repo) *Service {
	return &Service{repo}
}

func (s *Service) Login(username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	if *user.Password != password {
		return nil, errors.New("Password is incorrect")
	}

	return user, nil
}
