// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countCategory = `-- name: CountCategory :one
select
    count(*)
from
    categories
where
    categories.name ilike $1
`

func (q *Queries) CountCategory(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, countCategory, name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countInventoryItems = `-- name: CountInventoryItems :one
select 
    count(*)
from 
    branch_items
join
    items
on
    branch_items.item_id = items.id
join
    categories
on
    items.category_id = categories.id
join
    branches
on
    branch_items.branch_id = branches.id
where
    categories.id ilike $1
and
    branches.id ilike $2
`

type CountInventoryItemsParams struct {
	ID   string
	ID_2 string
}

func (q *Queries) CountInventoryItems(ctx context.Context, arg CountInventoryItemsParams) (int64, error) {
	row := q.db.QueryRow(ctx, countInventoryItems, arg.ID, arg.ID_2)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createCategory = `-- name: CreateCategory :exec
insert into categories
(
    name,
    description,
    insert_date
)
values
(
    $1,
    $2,
    $3
)
`

type CreateCategoryParams struct {
	Name        string
	Description string
	InsertDate  pgtype.Timestamp
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.Exec(ctx, createCategory, arg.Name, arg.Description, arg.InsertDate)
	return err
}

const fetchCategory = `-- name: FetchCategory :many
select
    categories.id,
    categories.name,
    categories.description,
    categories.insert_date,
    categories.update_date
from
    categories
where
    categories.name ilike $1
order by
    categories.id
desc
limit $2
offset $3
`

type FetchCategoryParams struct {
	Name   string
	Limit  int32
	Offset int32
}

type FetchCategoryRow struct {
	ID          string
	Name        string
	Description string
	InsertDate  pgtype.Timestamp
	UpdateDate  pgtype.Timestamp
}

func (q *Queries) FetchCategory(ctx context.Context, arg FetchCategoryParams) ([]FetchCategoryRow, error) {
	rows, err := q.db.Query(ctx, fetchCategory, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchCategoryRow
	for rows.Next() {
		var i FetchCategoryRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.InsertDate,
			&i.UpdateDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchInventoryItems = `-- name: FetchInventoryItems :many
select 
    branch_items.id,
    items.name as item_name,
    branch_items.qty,
    categories.name as category_name,
    branches.name as branch_name,
    branch_items.position_code
from 
    branch_items
join
    items
on
    branch_items.item_id = items.id
join
    categories
on
    items.category_id = categories.id
join
    branches
on
    branch_items.branch_id = branches.id
where
    categories.id ilike $1
and
    branches.id ilike $2
order by 
    items.name
desc
limit $3
offset $4
`

type FetchInventoryItemsParams struct {
	ID     string
	ID_2   string
	Limit  int32
	Offset int32
}

type FetchInventoryItemsRow struct {
	ID           string
	ItemName     string
	Qty          int32
	CategoryName string
	BranchName   string
	PositionCode pgtype.Text
}

func (q *Queries) FetchInventoryItems(ctx context.Context, arg FetchInventoryItemsParams) ([]FetchInventoryItemsRow, error) {
	rows, err := q.db.Query(ctx, fetchInventoryItems,
		arg.ID,
		arg.ID_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FetchInventoryItemsRow
	for rows.Next() {
		var i FetchInventoryItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.ItemName,
			&i.Qty,
			&i.CategoryName,
			&i.BranchName,
			&i.PositionCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
