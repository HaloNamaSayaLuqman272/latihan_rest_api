package products

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/ai"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
)

type Service interface {
	CreateProduct(ctx context.Context, productReq *ProductRequest) (models.Product, error)
	GetProductByID(ctx context.Context, id uint) (models.Product, error)
	GetProductsByCategory(ctx context.Context, pagination utils.Pagination, categoryId uint) (utils.Pagination, error)
	GetAllProducts(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	GetProductsByRecommendationAI(ctx context.Context, pagination utils.Pagination, promptReq ai.ProductRecommendationRequest) (utils.Pagination, error)
	UpdateProduct(ctx context.Context, productReq *ProductRequest, id uint) (models.Product, error)
	DeleteProduct(ctx context.Context, id uint) error
}

type service struct {
	create
	getbyid
	getbycategory
	getall
	getbyrecommendation
	update
	delete
}

func New(repository *gorm.DB) Service {
	return service{
		create:              create{repository: repository},
		getbyid:             getbyid{repository: repository},
		getbycategory:       getbycategory{repository: repository},
		getall:              getall{repository: repository},
		getbyrecommendation: getbyrecommendation{repository: repository},
		update:              update{repository: repository, getbyid: getbyid{repository: repository}},
		delete:              delete{repository: repository, getbyid: getbyid{repository: repository}},
	}
}
