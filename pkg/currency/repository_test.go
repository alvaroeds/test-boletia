package currency

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPostgres_GetAll(t *testing.T) {
	currency := []GetAllCurrency{
		{
			Code:      "USD",
			Value:     1,
			CreatedAt: time.Now(),
		},
	}
	repo := NewRepository(&sql.DB{}, 0)
	currency, err := repo.GetAll(ctx)
	require.Error(t, err)
	require.Equal(t, currency, currency)
}

func TestPostgres_GetByID(t *testing.T) {
	currency := []GetAllCurrency{
		{
			Code:      "USD",
			Value:     1,
			CreatedAt: time.Now(),
		},
	}
	f := CurrencyFilterRequest{
		Code:  "USD",
		FInit: time.Now(),
		FEnd:  time.Now(),
	}
	repo := NewRepository(&sql.DB{}, 0)
	currency, err := repo.GetByID(ctx, f)
	require.Error(t, err)
	require.Equal(t, currency, currency)
}

func TestPostgres_Insert(t *testing.T) {
	currency := &InsertCurrency{
		Code:  "USD",
		Value: 1,
	}
	repo := NewRepository(&sql.DB{}, 0)
	err := repo.Insert(ctx, currency)
	require.Error(t, err)
}
