package models

import (
	"database/sql/driver"
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
	Role           Role           `json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

type Role string

const (
	Admin   Role = "admin"
	Enduser Role = "user"
)

func (p *Role) Scan(value any) error {
	*p = Role(value.([]byte))
	return nil
}

func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}
