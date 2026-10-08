package users

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	get
}

func (u update) UpdateProfile(ctx context.Context, editProfileRequest *EditProfileRequest, id uint) (User, error) {
	user := models.User{
		Username:       editProfileRequest.Username,
		Email:          editProfileRequest.Email,
		Password:       editProfileRequest.Password,
		PhoneNumber:    editProfileRequest.PhoneNumber,
		Address:        editProfileRequest.Address,
		ProvinceID:     editProfileRequest.ProvinceID,
		DistrictID:     editProfileRequest.DistrictID,
		ProfilePicture: editProfileRequest.ProfilePicture,
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&user)
	if err := result.Error; err != nil {
		return User{}, err
	}

	record, err := u.GetProfile(ctx, id)
	if err != nil {
		return User{}, err
	}

	return record, nil
}
