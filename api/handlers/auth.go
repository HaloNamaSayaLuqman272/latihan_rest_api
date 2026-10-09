package handlers

import (
	"latihan_rest_api/api/middlewares"
	"latihan_rest_api/auth"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/dtos"
	"latihan_rest_api/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Auth struct {
	auth      auth.Service
	jwtConfig middlewares.JWTConfig
}

func (a Auth) RegisterUser(ctx *echo.Context) error {
	registerRequest := new(auth.RegisterRequest)
	if err := ctx.Bind(registerRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(registerRequest); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    utils.GetValidationErrorMessage(err.Error()),
		})
	}

	user, err := a.auth.RegisterUser(ctx.Request().Context(), registerRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "user registration failed",
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.Response[models.User]{
		Status:  "success",
		Message: "user registrated",
		Data:    user,
	})
}

func (a Auth) LoginUser(ctx *echo.Context) error {
	loginRequest := new(auth.LoginRequest)
	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(loginRequest); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    utils.GetValidationErrorMessage(err.Error()),
		})
	}

	user, err := a.auth.LoginUser(ctx.Request().Context(), loginRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "user login failed",
		})
	}

	token, err := a.jwtConfig.GenerateToken(user.ID, models.Role(user.Role))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "token generate failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}

func NewAuth(auth auth.Service, jwtConfig middlewares.JWTConfig) Auth {
	authHandler := Auth{
		auth:      auth,
		jwtConfig: jwtConfig,
	}

	return authHandler
}
