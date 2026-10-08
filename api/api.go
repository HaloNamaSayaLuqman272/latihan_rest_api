package api

import (
	"latihan_rest_api/auth"
	"latihan_rest_api/users"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB) *echo.Echo {
	var (
		e           = echo.New()
		authService = auth.New(repository)
		userService = users.New(repository)
		authHandler = 
	)
}
