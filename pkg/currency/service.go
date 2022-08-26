package currency

import (
	"context"
	"database/sql"
)

type Service interface {
	GetByCurrency(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error)
	Insert(ctx context.Context, req *InsertCurrency) error
}

type service struct {
	repo Repository
}

func NewService(db *sql.DB) Service {
	return &service{
		repo: NewRepository(db),
	}
}

func (s service) GetByCurrency(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error) {

	if f.Code == "all" {
		return s.repo.GetAll(ctx)
	}

	return s.repo.GetByID(ctx, f)
}

func (s service) Insert(ctx context.Context, req *InsertCurrency) error {
	return s.repo.Insert(ctx, req)
}
