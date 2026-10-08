package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	Password       string         `json:"-"`
	PhoneNumber    string         `json:"phone_number"`
	Address        string         `json:"address"`
	ProvinceID     uint           `json:"province_id"`
	CityID         uint           `json:"city_id"`
	DistrictID     uint           `json:"district_id"`
	ProfilePicture string         `json:"profile_picture"`
	Role           string         `json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

type RegisterRequest struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
	PhoneNumber string `json:"phone_number" validate:"required,containsNumber"`
	Address     string `json:"address" validate:"required,containsNumber"`
	ProvinceID  uint   `json:"province_id" validate:"required"`
	CityID      uint   `json:"city_id" validate:"required"`
	DistrictID  uint   `json:"district_id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
