package controllers

import (
	"encoding/json"
	"log"

	"github.com/syamsv/apollo/api/db"
	"github.com/syamsv/apollo/api/schema"
	"github.com/syamsv/apollo/api/session"
	"github.com/syamsv/apollo/pkg/mailer"
	"github.com/syamsv/apollo/pkg/models"
	"github.com/syamsv/apollo/pkg/template"
	"golang.org/x/crypto/bcrypt"
)

func CacheUser(user *models.Users) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(password)
	jsondata, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	id, err := session.StoreUserDetials(string(jsondata))
	if err != nil {
		return "", err
	}
	go func() {
		if err := mailer.SendActivactionMail(user.Email, template.ReturnHtmlTemplate(id)); err != nil {
			log.Println(err)
		}
	}()
	return id, nil
}

func VerifyUser(loginCreds *schema.LoginCreds) (string, error) {
	user, err := db.User.FetchProfileByEmail(loginCreds.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCreds.Password)); err != nil {

		return "bad password", err
	}

	sessionId, err := session.GenerateSession(user)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func ActivateUser(id string) error {
	user, err := session.GetUserDetails(id)
	if err != nil {
		return err
	}
	if _, err := db.User.CreateUser(user); err != nil {
		return err
	}
	return err
}
