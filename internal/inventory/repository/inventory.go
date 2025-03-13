package repository

import (
	"app/api/inventory/operation"
	database "app/database/main"
	"context"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
)

type Inventory interface {
	CountInventoryItems(ctx context.Context, params *operation.InventoryRequest) (int, int, error)
	FetchInventoryItems(ctx context.Context, params *operation.InventoryRequest) (*[]database.FetchInventoryItemsRow, error)
}

type InventoryService struct {
	Inventory
}

func InventoryRepository(inventory Inventory) *InventoryService {
	return &InventoryService{
		Inventory: inventory,
	}
}

func (r *Repository) CountInventoryItems(ctx context.Context, params *operation.InventoryRequest) (int, int, error) {
	args := database.CountInventoryItemsParams{
		CategoryID: pgtype.Text{
			String: fmt.Sprintf("%%%v%%", params.Category),
			Valid:  true,
		},
		LocationID: pgtype.Text{
			String: fmt.Sprintf("%%%v%%", params.LocationId),
			Valid:  true,
		},
	}

	total, err := r.read.CountInventoryItems(ctx, args)
	if err != nil {
		return 0, 0, err
	}

	page := math.Ceil(float64(total) / float64(params.Limit))

	return int(total), int(page), nil
}

func (r *Repository) FetchInventoryItems(ctx context.Context, params *operation.InventoryRequest) (*[]database.FetchInventoryItemsRow, error) {
	args := database.FetchInventoryItemsParams{
		CategoryID: pgtype.Text{
			String: fmt.Sprintf("%%%v%%", params.Category),
			Valid:  true,
		},
		LocationID: pgtype.Text{
			String: fmt.Sprintf("%%%v%%", params.LocationId),
			Valid:  true,
		},
		Limit:  int32(params.Limit),
		Offset: (int32(params.Page) - 1) * int32(params.Limit),
	}

	items, err := r.read.FetchInventoryItems(ctx, args)
	if err != nil {
		return nil, err
	}

	return &items, nil
}
