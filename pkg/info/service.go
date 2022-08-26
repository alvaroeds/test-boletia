package info

import (
	"context"
	"database/sql"
)

type Service interface {
	InsertRequest(ctx context.Context, req *InsertRequest) error
}

type service struct {
	repo Repository
}

func NewService(db *sql.DB) Service {
	return &service{
		repo: NewRepository(db),
	}
}

func (s service) InsertRequest(ctx context.Context, req *InsertRequest) error {
	return s.repo.InsertRequest(ctx, req)
}
