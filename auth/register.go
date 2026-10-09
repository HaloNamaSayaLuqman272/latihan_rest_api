package auth

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
)

type register struct {
	repository *gorm.DB
}

func (r register) RegisterUser(ctx context.Context, registerRequest *RegisterRequest) (models.User, error) {
	// ini adalah bagian password di-generate menjadi token bcrypt
	password, err := utils.GeneratePassword(registerRequest.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username:    registerRequest.Username,
		Email:       registerRequest.Email,
		Password:    string(password),
		PhoneNumber: registerRequest.PhoneNumber,
		Address:     registerRequest.Address,
		ProvinceID:  registerRequest.ProvinceID,
		CityID:      registerRequest.CityID,
		DistrictID:  registerRequest.DistrictID,
		Role:        models.Enduser,
	}
	if err := r.repository.WithContext(ctx).Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
