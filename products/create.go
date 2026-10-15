package products

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type create struct {
	repository *gorm.DB
}

func (c create) CreateProduct(ctx context.Context, productReq *ProductRequest) (models.Product, error) {
	product := models.Product{
		ProductName: productReq.ProductName,
		CategoryID:  productReq.CategoryID,
		Category:    productReq.Category,
		Description: productReq.Description,
		Company:     productReq.Company,
		Barcode:     productReq.Barcode,
		Weight:      productReq.Weight,
		ExpiredDate: productReq.ExpiredDate,
		Price:       productReq.Price,
		Stock:       productReq.Stock,
		ImageLink:   productReq.ImageLink,
	}

	result := c.repository.WithContext(ctx).Create(&product)
	if err := result.Error; err != nil {
		return models.Product{}, err
	}

	record := new(models.Product)

	if err := result.WithContext(ctx).Preload("Category").Last(record).Error; err != nil {
		return models.Product{}, err
	}

	return *record, nil
}
