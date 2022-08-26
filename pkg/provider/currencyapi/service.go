package currencyapi

import (
	"context"
	"github.com/alvaroeds/test-boletia/pkg/request"
)

type Service interface {
	GetCurrencys(ctx context.Context) (*Latest, *QueryInfo, error)
	GetStatus(ctx context.Context) (*Status, error)
}

type service struct {
	repo Repository
}

func (s service) GetCurrencys(ctx context.Context) (*Latest, *QueryInfo, error) {
	return s.repo.GetCurrencys(ctx)
}

func (s service) GetStatus(ctx context.Context) (*Status, error) {
	return s.repo.GetStatus(ctx)
}

func NewService(req request.Request, token string) Service {
	return service{
		repo: newRepository(req, token),
	}
}
