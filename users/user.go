package users

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `json:"id" form:"-" gorm:"primaryKey"`
	SessionID      string         `json:"session_id"`
	Username       string         `json:"username" gorm:"uniqueIndex;not null"`
	Email          string         `json:"email" gorm:"uniqueIndex;not null"`
	Password       string         `json:"password"`
	PhoneNumber    string         `json:"phone_number"`
	Address        string         `json:"address" gorm:"type:text"`
	ProvinceID     uint           `json:"province_id"`
	CityID         uint           `json:"city_id"`
	DistrictID     uint           `json:"district_id"`
	ProfilePicture string         `json:"profile_picture"`
	Role           string         `json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

type EditProfileRequest struct {
	Username       string `form:"username" validate:"required"`
	Email          string `form:"email" validate:"required,email"`
	Password       string `form:"password" validate:"min=8,containsSpecialCharacter,containsNumber"`
	PhoneNumber    string `form:"phone_number" validate:"required,min=10,containsNumberOnly"`
	Address        string `form:"address" validate:"required,containsNumber"`
	ProvinceID     uint   `form:"province_id" validate:"required"`
	DistrictID     uint   `form:"district_id" validate:"required"`
	ProfilePicture string
	File           *multipart.FileHeader
}
