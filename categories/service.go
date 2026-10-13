package categories

import (
	"context"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/utils"

	"gorm.io/gorm"
)

type Service interface {
	CreateCategory(ctx context.Context, categoryReq *CategoryRequest) (models.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (models.Category, error)
	GetAllCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	UpdateCategory(ctx context.Context, categoryReq *CategoryRequest, id uint) (models.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}

type service struct {
	create
	getbyid
	getall
	update
	delete
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		create:  create{repository: repository},
		getbyid: getbyid{repository: repository},
		getall:  getall{repository: repository},
		update:  update{repository: repository, getbyid: getbyid{repository: repository}},
		delete:  delete{repository: repository, getbyid: getbyid{repository: repository}},
	}
}
