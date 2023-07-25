package handler

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/syamsv/apollo/api/controllers"
	"github.com/syamsv/apollo/api/schema"
	"github.com/syamsv/apollo/api/views"
	"github.com/syamsv/apollo/pkg/models"
	"gorm.io/gorm"
)

func AuthLogin(c *fiber.Ctx) error {
	loginCreds := new(schema.LoginCreds)
	if err := c.BodyParser(loginCreds); err != nil {
		return views.InvalidParams(c)
	}

	if err := validator.New().Struct(loginCreds); err != nil {
		return views.BadRequest(c)
	}
	accesstoken, err := controllers.VerifyUser(loginCreds)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return views.BadRequest(c)
		}
		if accesstoken == "bad password" {
			return views.Unauthorized(c)
		}
		return views.InternalServerError(c, err)
	}

	return views.SuccessResponse(c, "Successfully	 logged in")
}

func AuthRegister(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return views.InvalidParams(c)
	}
	if err := validator.New().Struct(user); err != nil {
		return views.BadRequest(c)
	}
	if _, err := controllers.CacheUser(user); err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return views.ErrorResponse(c, 400, "Email already in use with another account")
		}
		return views.InternalServerError(c, err)
	}
	return views.CreatedResponse(c, "Successfully registered")
}

func AuthActivateAccount(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return views.BadRequest(c)
	}
	if err := controllers.ActivateUser(id); err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return views.ErrorResponse(c, 400, "Email already in use with another account")
		}
		return views.InternalServerError(c, err)
	}

	return views.SuccessResponse(c, "Successsfully activated the account")
}
