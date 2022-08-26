package info

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgres_InsertQuery(t *testing.T) {
	query := &InsertRequest{
		Method: "GET",
		Path:   "/currency/usd",
		Code:   200,
		Time:   0.9312,
	}
	repo := NewRepository(&sql.DB{}, 0)
	err := repo.InsertRequest(ctx, query)
	require.Error(t, err)
}
