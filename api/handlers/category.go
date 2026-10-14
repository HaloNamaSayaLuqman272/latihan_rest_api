package handlers

import (
	"latihan_rest_api/categories"
	"latihan_rest_api/db/models"
	"latihan_rest_api/pkg/dtos"
	"latihan_rest_api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Categories struct {
	categories categories.Service
}

func (c Categories) CreateCategory(ctx *echo.Context) error {
	categoryReq := ctx.Get("validatedBody").(*categories.CategoryRequest)

	category, err := c.categories.CreateCategory(ctx.Request().Context(), categoryReq)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "create category failed",
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.Response[models.Category]{
		Status:  "success",
		Message: "category created",
		Data:    category,
	})
}

func (c Categories) GetCategoryByID(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	category, err := c.categories.GetCategoryByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "category not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[models.Category]{
		Status:  "success",
		Message: "category found",
		Data:    category,
	})
}

func (c Categories) GetAllCategories(ctx *echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	sort := ctx.QueryParam("sort")
	search := ctx.QueryParam("search")

	pagination := utils.Pagination{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		Search:  search,
		Keyword: "category_name",
	}

	categoriesData, err := c.categories.GetAllCategories(ctx.Request().Context(), pagination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch categories failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[utils.Pagination]{
		Status:  "success",
		Message: "all categories",
		Data:    categoriesData,
	})
}

func (c Categories) UpdateCategory(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	categoryReq := ctx.Get("validatedBody").(*categories.CategoryRequest)

	category, err := c.categories.UpdateCategory(ctx.Request().Context(), categoryReq, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "update category failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[models.Category]{
		Status:  "success",
		Message: "category updated",
		Data:    category,
	})
}

func (c Categories) DeleteCategory(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	err = c.categories.DeleteCategory(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "delete category failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[any]{
		Status:  "success",
		Message: "category deleted",
	})
}

func NewCategories(categories categories.Service) Categories {
	return Categories{
		categories: categories,
	}
}
