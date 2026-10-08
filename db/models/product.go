package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ProductName string         `json:"product_name"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `json:"category"`
	Description string         `json:"description"`
	Company     string         `json:"company"`
	Barcode     string         `json:"barcode"`
	Weight      uint           `json:"weight"`
	ExpiredDate string         `json:"expired_date"`
	Price       float64        `json:"price"`
	Stock       uint           `json:"stock"`
	ImageLink   string         `json:"image_link"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
