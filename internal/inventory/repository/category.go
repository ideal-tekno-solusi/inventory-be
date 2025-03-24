package repository

import (
	database "app/database/main"
	"context"
	"fmt"
	"math"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Category interface {
	CountCategory(ctx context.Context, name string, limit int) (int, int, error)
	FetchCategory(ctx context.Context, name string, page, limit int) (*[]database.FetchCategoryRow, error)
	CreateCategory(ctx context.Context, name, description string) error
	UpdateCategory(ctx context.Context, id, name, description string) error
	DeleteCategory(ctx context.Context, id string) error
}

type CategoryService struct {
	Category
}

func CategoryRepository(category Category) *CategoryService {
	return &CategoryService{
		Category: category,
	}
}

func (r *Repository) CountCategory(ctx context.Context, name string, limit int) (int, int, error) {
	args := fmt.Sprintf("%%%v%%", name)

	total, err := r.read.CountCategory(ctx, args)
	if err != nil {
		return 0, 0, err
	}

	page := math.Ceil(float64(total) / float64(limit))

	return int(total), int(page), nil
}

func (r *Repository) FetchCategory(ctx context.Context, name string, page, limit int) (*[]database.FetchCategoryRow, error) {
	args := database.FetchCategoryParams{
		Name:   fmt.Sprintf("%%%v%%", name),
		Limit:  int32(limit),
		Offset: (int32(page) - 1) * int32(limit),
	}

	categories, err := r.read.FetchCategory(ctx, args)
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (r *Repository) CreateCategory(ctx context.Context, name, description string) error {
	args := database.CreateCategoryParams{
		Name:        name,
		Description: description,
		InsertDate: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}

	err := r.write.CreateCategory(ctx, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateCategory(ctx context.Context, id, name, description string) error {
	args := database.UpdateCategoryParams{
		Name:        name,
		Description: description,
		UpdateDate: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ID: id,
	}

	err := r.write.UpdateCategory(ctx, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCategory(ctx context.Context, id string) error {
	args := database.DeleteCategoryParams{
		DeleteDate: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ID: id,
	}

	err := r.write.DeleteCategory(ctx, args)
	if err != nil {
		return err
	}

	return nil
}
