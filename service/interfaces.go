package service

import "user-service/models"

type UserStore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(username string) (*models.User, error)
}
