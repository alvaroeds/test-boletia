package currency

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPostgres_GetAll(t *testing.T) {
	ctx := context.Background()
	currency := []GetAllCurrency{
		{
			Code:      "USD",
			Value:     1,
			CreatedAt: time.Now(),
		},
	}
	repo := NewRepository(&sql.DB{})
	currency, err := repo.GetAll(ctx)
	require.Error(t, err)
	require.Equal(t, currency, currency)
}
