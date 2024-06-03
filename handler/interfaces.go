package handler

import "user-service/models"

type UserService interface {
	CreateUser(user *models.User) error
	LoginUser(email, password string) (*models.User, error)
	GetUserProfile(username string) (*models.User, error)
}
