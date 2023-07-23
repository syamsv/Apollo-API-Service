package controllers

import (
	"encoding/json"

	"github.com/syamsv/apollo/api/db"
	"github.com/syamsv/apollo/api/jwt"
	"github.com/syamsv/apollo/api/schema"
	"github.com/syamsv/apollo/api/session"
	"github.com/syamsv/apollo/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func CacheUser(user *models.Users) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(password)
	if user.Role == "" {
		user.Role = "user"
	}
	jsondata, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	id, err := session.StoreUserDetials(string(jsondata))
	if err != nil {
		return "", err
	}
	return id, nil
}

func CreateUser(user *models.Users) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(password)
	if user.Role == "" {
		user.Role = "user"
	}
	if user, err = db.User.CreateUser(user); err != nil {
		return "", err
	}
	return user.ID.String(), err
}

func VerifyUser(loginCreds *schema.LoginCreds) (string, string, error) {
	user, err := db.User.FetchProfileByEmail(loginCreds.Email)
	if err != nil {
		return "", "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCreds.Password)); err != nil {

		return "bad password", "", err
	}

	accesstoken, refreshtoken, err := jwt.GenerateTokens(user)
	if err != nil {
		return "", "", err
	}
	return accesstoken, refreshtoken, err

}
