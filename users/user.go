package users

import (
	"mime/multipart"
)

type EditProfileRequest struct {
	Username       string `form:"username" validate:"required"`
	Email          string `form:"email" validate:"required,email"`
	Password       string `form:"password" validate:"min=8,containsSpecialCharacter,containsNumber"`
	PhoneNumber    string `form:"phone_number" validate:"required,min=10,containsNumberOnly"`
	Address        string `form:"address" validate:"required,containsNumber"`
	ProvinceID     uint   `form:"province_id" validate:"required"`
	CityID         uint   `form:"city_id" validate:"required"`
	DistrictID     uint   `form:"district_id" validate:"required"`
	ProfilePicture string
	File           *multipart.FileHeader
}
