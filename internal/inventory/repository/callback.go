package repository

import (
	database "app/database/main"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Callback interface {
	GetChallenge(ctx context.Context, codeChallenge string) (*database.Challenge, error)
	DeleteChallenge(ctx context.Context, codeChallenge string) error
}

type CallbackService struct {
	Callback
}

func CallbackRepository(callback Callback) *CallbackService {
	return &CallbackService{
		Callback: callback,
	}
}

func (r *Repository) GetChallenge(ctx context.Context, codeChallenge string) (*database.Challenge, error) {
	args := pgtype.Text{
		String: codeChallenge,
		Valid:  true,
	}

	data, err := r.read.GetChallenge(ctx, args)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Repository) DeleteChallenge(ctx context.Context, codeChallenge string) error {
	args := pgtype.Text{
		String: codeChallenge,
		Valid:  true,
	}

	err := r.write.DeleteChallenge(ctx, args)
	if err != nil {
		return err
	}

	return nil
}
