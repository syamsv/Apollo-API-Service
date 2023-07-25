package session

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/syamsv/apollo/pkg/models"
)

type Session struct {
	Email string    `json:"email"`
	ID    uuid.UUID `json:"id"`
	Role  string    `json:"role"`
}

func StoreUserDetials(data string) (string, error) {
	uid := uuid.NewString()
	if err := manager.Verify.SetValue(uid, data, 15*time.Minute); err != nil {
		return "", err
	}
	return uid, nil
}

func GetUserDetails(uid string) (*models.Users, error) {
	userdetials, err := manager.Verify.GetValue(uid)
	if err != nil {
		return nil, err
	}
	user := new(models.Users)
	if err := json.Unmarshal([]byte(userdetials), user); err != nil {
		return nil, err
	}

	return user, nil
}

func GenerateSession(user *models.Users) (string, error) {
	session := uuid.NewString()
	s := &Session{
		Email: user.Email,
		ID:    user.ID,
	}
	jsondata, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}

	if err := manager.Auth.SetValue(session, string(jsondata), time.Hour); err != nil {
		return "", err
	}
	return session, nil

}

func GetSession(session string) (*Session, error) {
	data, err := manager.Auth.GetValue(session)
	if err != nil {
		log.Println(err)
	}
	s := new(Session)
	if err := json.Unmarshal([]byte(data), s); err != nil {
		return nil, err
	}
	return s, nil
}
