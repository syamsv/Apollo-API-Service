package users

import (
	"github.com/syamsv/apollo/pkg/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Interface {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateUser(user *models.Users) (*models.Users, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) FetchProfileByEmail(email string) (*models.Users, error) {

	user := new(models.Users)

	if err := r.DB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
