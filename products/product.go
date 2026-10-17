package products

import (
	"latihan_rest_api/db/models"
	"mime/multipart"
)

type ProductRequest struct {
	ProductName string          `form:"product_name" validate:"required"`
	CategoryID  uint            `form:"category_id" validate:"required"`
	Category    models.Category `form:"category" validate:"required"`
	Description string          `form:"description" validate:"required"`
	Company     string          `form:"company" validate:"required"`
	Barcode     string          `form:"barcode"`
	Weight      uint            `form:"weight" validate:"required"`
	ExpiredDate string          `form:"expired_date"`
	Price       float64         `form:"price" validate:"required"`
	Stock       uint            `form:"stock" validate:"required"`
	ImageLink   string
	File        *multipart.FileHeader
}
