package handlers

import (
	"latihan_rest_api/api/middlewares"
	"latihan_rest_api/pkg/dtos"
	"latihan_rest_api/pkg/fileupload"
	"latihan_rest_api/pkg/utils"
	"latihan_rest_api/users"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v5"
)

type Users struct {
	users    users.Service
	uploader fileupload.Uploader
}

func (u Users) GetProfile(ctx *echo.Context) error {
	userID, err := middlewares.GetUserID(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid token",
		})
	}

	user, err := u.users.GetProfile(ctx.Request().Context(), uint(userID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[any]{
		Status:  "success",
		Message: "user data found",
		Data:    user,
	})
}

func (u Users) UpdateProfile(ctx *echo.Context) error {
	userID, err := middlewares.GetUserID(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid token",
		})
	}

	editReq := new(users.EditProfileRequest)
	if err := ctx.Bind(editReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(editReq); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    utils.GetValidationErrorMessage(err.Error()),
		})
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "file not found",
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	isFileValid := utils.ValidateFile(ext)
	if !isFileValid {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid format file",
		})
	}

	fileLink, err := u.uploader.UploadFile(ctx.Request().Context(), file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "update file failed",
		})
	}
	editReq.ProfilePicture = fileLink

	user, err := u.users.UpdateProfile(ctx.Request().Context(), editReq, uint(userID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "update profile failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[users.User]{
		Status:  "success",
		Message: "updated profile",
		Data:    user,
	})
}

func NewUsers(user users.Service, uploader fileupload.Uploader) Users {
	usersHandler := Users{
		users:    user,
		uploader: uploader,
	}

	return usersHandler
}
