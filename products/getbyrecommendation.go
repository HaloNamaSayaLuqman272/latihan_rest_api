package products

import (
	"context"
	"encoding/json"
	"errors"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/ai"
	"latihan_rest_api/pkg/utils"
	"log"
	"math"

	"gorm.io/gorm"
)

type getbyrecommendation struct {
	repository *gorm.DB
}

func (g getbyrecommendation) GetProductsByRecommendationAI(ctx context.Context, pagination utils.Pagination, promptReq ai.ProductRecommendationRequest) (utils.Pagination, error) {
	products := []models.Product{}
	if err := g.repository.WithContext(ctx).Find(&products).Error; err != nil {
		return utils.Pagination{}, err
	}
	if len(products) == 0 {
		return utils.Pagination{}, errors.New("no products available")
	}

	aiProducts := make([]ai.Product, len(products))
	for i, p := range products {
		aiProducts[i] = ai.Product{
			ID:          p.ID,
			ProductName: p.ProductName,
			Description: p.Description,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
		}
	}
	promptReq.Products = aiProducts

	aiService := ai.InitService()
	res, err := aiService.GetProductRecommendation(ctx, promptReq)
	if err != nil {
		return utils.Pagination{}, err
	}

	resBody := res.Choices[0].Message.Content
	log.Println("DEBUG AI raw response:", resBody)
	var selected []struct {
		ID uint `json:"id"`
	}
	if err := json.Unmarshal([]byte(resBody), &selected); err != nil {
		return utils.Pagination{}, err
	}
	if len(selected) == 0 {
		return utils.Pagination{}, errors.New("no products selected AI")
	}

	ids := make([]uint, len(selected))
	for i, s := range selected {
		ids[i] = s.ID
	}

	var totalRows int64
	if err := g.repository.Model(&models.Product{}).Where("id IN ?", ids).Count(&totalRows).Error; err != nil {
		return utils.Pagination{}, err
	}
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	var recommendedProducts []models.Product
	if err := g.repository.WithContext(ctx).Where("id IN ?", ids).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Preload("Category").Find(&recommendedProducts).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = recommendedProducts
	return pagination, nil
}
