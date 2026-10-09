package users

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	get
}

func (u update) UpdateProfile(ctx context.Context, editProfileRequest *EditProfileRequest, id uuid.UUID) (models.User, error) {
	password, err := utils.GeneratePassword(editProfileRequest.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username:       editProfileRequest.Username,
		Email:          editProfileRequest.Email,
		Password:       string(password),
		PhoneNumber:    editProfileRequest.PhoneNumber,
		Address:        editProfileRequest.Address,
		ProvinceID:     editProfileRequest.ProvinceID,
		CityID:         editProfileRequest.CityID,
		DistrictID:     editProfileRequest.DistrictID,
		ProfilePicture: editProfileRequest.ProfilePicture,
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&user)
	if err := result.Error; err != nil {
		return models.User{}, err
	}

	record, err := u.GetProfile(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return record, nil
}
