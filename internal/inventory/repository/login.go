package repository

import (
	database "app/database/main"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Login interface {
	CreateChallenge(ctx context.Context, codeVerifier, codeChallenge, codeChallengeMethod string) error
}

type LoginService struct {
	Login
}

func LoginRepository(login Login) *LoginService {
	return &LoginService{
		Login: login,
	}
}

func (r *Repository) CreateChallenge(ctx context.Context, codeVerifier, codeChallenge, codeChallengeMethod string) error {
	args := database.CreateChallengeParams{
		CodeVerifier: pgtype.Text{
			String: codeVerifier,
			Valid:  true,
		},
		CodeChallenge: pgtype.Text{
			String: codeChallenge,
			Valid:  true,
		},
		CodeChallengeMethod: pgtype.Text{
			String: codeChallengeMethod,
			Valid:  true,
		},
	}

	err := r.write.CreateChallenge(ctx, args)
	if err != nil {
		return err
	}

	return nil
}
