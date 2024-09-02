package auth

import (
	"errors"

	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/third_party/argon2"
)

func NewService(repo *Repo, argon2 *argon2.Argon2) *Service {
	return &Service{repo, argon2}
}

func (s *Service) Login(username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	isMatch, err := s.argon2.CompareHashPassword(password, *user.Password)
	if err != nil {
		return nil, err
	}

	if !isMatch {
		return nil, errors.New("Password is incorrect")
	}

	return user, nil
}
