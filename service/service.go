package service

import (
	"user-service/errors"
	"user-service/models"
)

type service struct {
	userStore UserStore
}

func New(userStore UserStore) service {
	return service{userStore: userStore}
}

func (s *service) CreateUser(user *models.User) error {
	existingUser, _ := s.userStore.GetUserByEmail(user.Email)
	if existingUser != nil {
		return &errors.BadRequest{Message: "userEmail Already Exists"}
	}

	err := s.userStore.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) LoginUser(email, password string) (*models.User, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil || user.Password != password {
		return nil, &errors.ValidationError{Message: "password is not matching with your email"}
	}
	return user, nil
}

func (s *service) GetUserProfile(email string) (*models.User, error) {
	return s.userStore.GetUserByEmail(email)
}
