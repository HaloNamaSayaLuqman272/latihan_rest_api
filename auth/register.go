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

func (r register) RegisterUser(ctx context.Context, registerRequest *RegisterRequest) (User, error) {
	// ini adalah bagian password di-generate menjadi token bcrypt
	password, err := utils.GeneratePassword(registerRequest.Password)
	if err != nil {
		return User{}, err
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
	result := r.repository.WithContext(ctx).Create(&user)
	if err := result.Error; err != nil {
		return User{}, err
	}

	record := new(User)

	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		// ini adalah proses data user baru disimpan ke dalam database
		return User{}, err
	}

	return *record, nil
}
