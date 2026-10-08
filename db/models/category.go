package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryName string         `json:"category_name" gorm:"uniqueIndex:idx_name_category_active,where:deleted_at IS NULL"`
	Description  string         `json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
