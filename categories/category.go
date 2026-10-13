package categories

type CategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
}
