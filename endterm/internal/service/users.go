package service

import (
	"twitter/internal/model"
	"twitter/internal/store"
)

type UsersService interface {
	CreateUser(user *model.User) (int, error)
	GetUser(username string) (*model.User, error)
}

type service struct {
	store store.Store
}

func NewUserService(s store.Store) UsersService {
	return &service{
		store: s,
	}
}

func (s *service) CreateUser(user *model.User) (int, error) {
	id, err := s.store.CreateUser(user)

	return id, err
}

func (s *service) GetUser(username string) (*model.User, error) {
	user, err := s.store.GetUser(username)

	return user, err
}
