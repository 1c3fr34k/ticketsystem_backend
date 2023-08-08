package users

import (
	"net/http"

	"github.com/1c3fr34k/ticketsystem_backend/database/models"
	"github.com/1c3fr34k/ticketsystem_backend/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (UsersRouter) CreateUser(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	hashedPassword, err := helpers.HashPassword(user.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	user.PasswordHash = hashedPassword

	db, _ := c.Get("db").(*gorm.DB)
	db.Create(&user)
	// if errors.Is(err, gorm.ErrDuplicatedKey) {
	// 	return c.JSON(http.StatusConflict, map[string]string{
	// 		"error": "User already exists",
	// 	})
	// } else if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"error": err.Error(),
	// 	})
	// }

	// }
	return c.JSON(http.StatusOK, user)
}

// GetUser godoc
//
//	@Summary		Get User.
//	@Description	gets a user by id.
//	@Tags			user
//	@Accept			*/*
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	models.User
//
// @Router       /accounts/{id} [get]
func (UsersRouter) GetUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	db, _ := c.Get("db").(*gorm.DB)
	db.First(&user, id)

	return c.JSON(http.StatusOK, user)
}

func (UsersRouter) GetUsers(c echo.Context) error {
	var users []models.User
	db, _ := c.Get("db").(*gorm.DB)
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}
