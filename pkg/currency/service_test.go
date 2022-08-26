package currency

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.Background()

type repositoryMock struct{}

func (r repositoryMock) GetByID(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error) {
	//TODO implement me
	panic("implement me")
}

func (r repositoryMock) GetAll(ctx context.Context) ([]GetAllCurrency, error) {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) Insert(ctx context.Context, req *InsertCurrency) error {
	// TODO implement me
	panic("implement me")
}

type GetAllCurrencyMock struct {
	repositoryMock
	GetAllCurrencyRes []GetAllCurrency
	GetAllCurrencyErr error
}

func (m GetAllCurrencyMock) GetAll(ctx context.Context) ([]GetAllCurrency, error) {
	return m.GetAllCurrencyRes, m.GetAllCurrencyErr
}

func (m GetAllCurrencyMock) GetByID(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error) {
	return m.GetAllCurrencyRes, m.GetAllCurrencyErr
}

func TestService_GetAll(t *testing.T) {
	allCurrency := []GetAllCurrency{
		{
			Code:  "USD",
			Value: 1,
		},
		{
			Code:  "TRY",
			Value: 17.958154,
		},
	}
	errorMock := errors.New("not found")
	tests := []struct {
		name              string
		filter            CurrencyFilterRequest
		getAllCurrencyRes []GetAllCurrency
		getAllCurrencyErr error
		GetAllRes         []GetAllCurrency
		GetAllErr         error
	}{
		{
			name: "success",
			filter: CurrencyFilterRequest{
				Code: "",
			},
			getAllCurrencyRes: allCurrency,
			getAllCurrencyErr: nil,
			GetAllRes:         allCurrency,
			GetAllErr:         nil,
		},
		{
			name: "all",
			filter: CurrencyFilterRequest{
				Code: "all",
			},
			getAllCurrencyRes: []GetAllCurrency{},
			getAllCurrencyErr: errorMock,
			GetAllRes:         []GetAllCurrency{},
			GetAllErr:         errorMock,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := GetAllCurrencyMock{
				GetAllCurrencyRes: test.getAllCurrencyRes,
				GetAllCurrencyErr: test.getAllCurrencyErr,
			}

			s := service{
				repo: repo,
			}
			resp, err := s.GetByCurrency(ctx, test.filter)
			assert.Equal(t, test.GetAllRes, resp)
			assert.Equal(t, test.GetAllErr, err)
		})
	}
}

type insertCurrency struct {
	repositoryMock
	InsertCurrencyErr error
}

func (m insertCurrency) Insert(ctx context.Context, req *InsertCurrency) error {
	return m.InsertCurrencyErr
}

func TestService_Insert(t *testing.T) {
	tests := []struct {
		name              string
		insertCurrencyReq *InsertCurrency
		insertCurrencyErr error
		insertErr         error
	}{
		{
			name: "success",
			insertCurrencyReq: &InsertCurrency{
				Code:  "USD",
				Value: 1,
			},
			insertCurrencyErr: nil,
			insertErr:         nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := insertCurrency{
				InsertCurrencyErr: test.insertCurrencyErr,
			}

			s := service{
				repo: repo,
			}
			err := s.Insert(ctx, test.insertCurrencyReq)
			assert.Equal(t, test.insertErr, err)
		})
	}
}
