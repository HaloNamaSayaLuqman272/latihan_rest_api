package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `json:"id" form:"-" gorm:"type:uuid;primaryKey"`
	Username       string         `json:"username" gorm:"uniqueIndex;not null"`
	Email          string         `json:"email" gorm:"uniqueIndex;not null"`
	Password       string         `json:"-"`
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	return nil
}

type Role string

const (
	Admin   Role = "admin"
	Enduser Role = "user"
)

func (p *Role) Scan(value any) error {
	switch v := value.(type) {
	case string:
		*p = Role(v)
	case []byte:
		*p = Role(v)
	default:
		return fmt.Errorf("unsupported role type: %T", value)
	}

	return nil
}

func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}
