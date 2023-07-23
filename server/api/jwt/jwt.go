package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/syamsv/apollo/config"
	"github.com/syamsv/apollo/pkg/models"
)

var (
	accessTokenSecretKey  = []byte(config.JWT_ACCESS_KEY_SECRET)
	refreshTokenSecretKey = []byte(config.JWT_REFRESH_KEY_SECRET)
)

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateTokens(users *models.Users) (string, string, error) {
	accesstoken, err := GenerateAccessToken(users.ID.String(), fmt.Sprintf("%s %s", users.FirstName, users.LastName))
	if err != nil {
		return "", "", err
	}
	refreshtoken, err := GenerateRefreshToken(users.ID.String(), fmt.Sprintf("%s %s", users.FirstName, users.LastName))
	if err != nil {
		return "", "", err
	}
	return accesstoken, refreshtoken, nil
}

func GenerateAccessToken(userID, username string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 2).Unix(),
			Issuer:    "Apollo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessTokenSecretKey)
}

func GenerateRefreshToken(userID, username string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "Apollo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenSecretKey)
}

func VerifyToken(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func ParseAccessToken(tokenString string) (*Claims, error) {
	return VerifyToken(tokenString, accessTokenSecretKey)
}

func ParseRefreshToken(tokenString string) (*Claims, error) {
	return VerifyToken(tokenString, refreshTokenSecretKey)
}
