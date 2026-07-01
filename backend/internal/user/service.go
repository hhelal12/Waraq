package user

import (
	md "backend/internal/domain"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}


func (s *Service) GetUserByID(id string) (*md.User, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// business logic example
	if user.Role == "" {
		user.Role = "user"
	}

	return user, nil
}


func (s *Service) GetAllUsers() ([]md.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// example business rule: filter or modify
	for i := range users {
		if users[i].Role == "" {
			users[i].Role = "user"
		}
	}

	return users, nil
}