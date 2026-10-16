package products

import (
	"context"
	"latihan_rest_api/db/models"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	getbyid
}

func (u update) UpdateProduct(ctx context.Context, productReq *ProductRequest, id uint) (models.Product, error) {
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

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&product)
	if err := result.Error; err != nil {
		return models.Product{}, err
	}

	record, err := u.getbyid.GetProductByID(ctx, id)
	if err != nil {
		return models.Product{}, err
	}

	return record, nil
}
