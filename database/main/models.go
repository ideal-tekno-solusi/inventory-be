// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Branch struct {
	ID          string
	Name        string
	Address     string
	Description pgtype.Text
	InsertDate  pgtype.Timestamp
	UpdateDate  pgtype.Timestamp
	DeleteDate  pgtype.Timestamp
}

type BranchItem struct {
	ID         string
	ItemID     pgtype.Text
	BranchID   pgtype.Text
	PositionID pgtype.Text
	Qty        int32
	InsertDate pgtype.Timestamp
	UpdateDate pgtype.Timestamp
	DeleteDate pgtype.Timestamp
}

type Category struct {
	ID          string
	Name        string
	Description string
	InsertDate  pgtype.Timestamp
	UpdateDate  pgtype.Timestamp
	DeleteDate  pgtype.Timestamp
}

type Item struct {
	ID         string
	CategoryID pgtype.Text
	Name       string
	InsertDate pgtype.Timestamp
	UpdateDate pgtype.Timestamp
	DeleteDate pgtype.Timestamp
}

type Position struct {
	ID         string
	Code       pgtype.Text
	BranchID   pgtype.Text
	InsertDate pgtype.Timestamp
	UpdateDate pgtype.Timestamp
	DeleteDate pgtype.Timestamp
}
