package products

import (
	"latihan_rest_api/db/models"
	"mime/multipart"
)

type ProductRequest struct {
	ProductName string          `json:"product_name" validate:"required"`
	CategoryID  uint            `json:"category_id" validate:"required"`
	Category    models.Category `json:"category" validate:"required"`
	Description string          `json:"description" validate:"required"`
	Company     string          `json:"company" validate:"required"`
	Barcode     string          `json:"barcode"`
	Weight      uint            `json:"weight" validate:"required"`
	ExpiredDate string          `json:"expired_date"`
	Price       float64         `json:"price" validate:"required"`
	Stock       uint            `json:"stock" validate:"required"`
	ImageLink   string
	File        *multipart.FileHeader
}
