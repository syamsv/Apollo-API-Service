package users

import "github.com/syamsv/apollo/pkg/models"

type Interface interface {
	FetchProfileByEmail(email string) (*models.Users, error)
	CreateUser(user *models.Users) (*models.Users, error)
}
