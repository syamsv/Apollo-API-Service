package users

import (
	"github.com/syamsv/apollo/pkg/models"
)

type userSvc struct {
	repo Interface
}

func NewService(r Interface) Interface {
	return &userSvc{repo: r}
}

func (s *userSvc) CreateUser(user *models.Users) (*models.Users, error) {
	return s.repo.CreateUser(user)
}

func (s *userSvc) FetchProfileByEmail(email string) (*models.Users, error) {
	return s.repo.FetchProfileByEmail(email)
}
