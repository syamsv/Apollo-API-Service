package handler

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo/api/controllers"
	"github.com/syamsv/apollo/api/jwt"
	"github.com/syamsv/apollo/api/schema"
	"github.com/syamsv/apollo/api/views"
	"github.com/syamsv/apollo/pkg/models"
	"gorm.io/gorm"
)

func extractTokenFromHeader(header string) string {
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}
	return ""
}

func AuthLogin(c *fiber.Ctx) error {
	loginCreds := new(schema.LoginCreds)
	if err := c.BodyParser(loginCreds); err != nil {
		return views.InvalidParams(c)
	}

	if err := validator.New().Struct(loginCreds); err != nil {
		return views.BadRequest(c)
	}
	accesstoken, refreshToken, err := controllers.VerifyUser(loginCreds)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return views.BadRequest(c)
		}
		if accesstoken == "bad password" {
			return views.Unauthorized(c)
		}
		return views.InternalServerError(c, err)
	}

	return views.SuccessResponse(c, &schema.JwtToken{
		AccessToken:  accesstoken,
		RefreshToken: refreshToken,
	})
}

func AuthRegister(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return views.InvalidParams(c)
	}
	if err := validator.New().Struct(user); err != nil {
		return views.BadRequest(c)
	}
	id, err := controllers.CacheUser(user)
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return views.ErrorResponse(c, 400, "Email already in use with another account")
		}
		return views.InternalServerError(c, err)
	}
	return views.CreatedResponse(c, fiber.Map{"id": id})
}

func AuthTokenVerify(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	token = extractTokenFromHeader(token)
	if token == "" {
		return views.Unauthorized(c)
	}

	_, err := jwt.ParseAccessToken(token)
	if err != nil {
		return views.Unauthorized(c)
	}
	return views.SuccessResponse(c, nil)
}

func AuthTokenRefresh(c *fiber.Ctx) error {
	refresh := c.Get("Authorization")
	refresh = extractTokenFromHeader(refresh)
	if refresh == "" {
		return views.Unauthorized(c)
	}

	refreshClaims, err := jwt.ParseRefreshToken(refresh)
	if err != nil {
		return views.Unauthorized(c)
	}
	accesstoken, err := jwt.GenerateAccessToken(refreshClaims.UserID, refreshClaims.Username)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.SuccessResponse(c, &schema.JwtToken{
		AccessToken:  accesstoken,
		RefreshToken: refresh,
	})
}
