package repository

import (
	"app/api/inventory/operation"
	database "app/database/main"
	"context"
	"fmt"
	"math"
)

type Category interface {
	CountCategory(ctx context.Context, params *operation.CategoryRequest) (int, int, error)
	FetchCategory(ctx context.Context, params *operation.CategoryRequest) (*[]database.FetchCategoryRow, error)
}

type CategoryService struct {
	Category
}

func CategoryRepository(category Category) *CategoryService {
	return &CategoryService{
		Category: category,
	}
}

func (r *Repository) CountCategory(ctx context.Context, params *operation.CategoryRequest) (int, int, error) {
	args := fmt.Sprintf("%%%v%%", params.Name)

	total, err := r.read.CountCategory(ctx, args)
	if err != nil {
		return 0, 0, err
	}

	page := math.Ceil(float64(total) / float64(params.Limit))

	return int(total), int(page), nil
}

func (r *Repository) FetchCategory(ctx context.Context, params *operation.CategoryRequest) (*[]database.FetchCategoryRow, error) {
	args := database.FetchCategoryParams{
		Name:   fmt.Sprintf("%%%v%%", params.Name),
		Limit:  int32(params.Limit),
		Offset: (int32(params.Page) - 1) * int32(params.Limit),
	}

	categories, err := r.read.FetchCategory(ctx, args)
	if err != nil {
		return nil, err
	}

	return &categories, nil
}
