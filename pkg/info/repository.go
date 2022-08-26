package info

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const timeout = time.Second * 5

type Repository interface {
	InsertRequest(ctx context.Context, req *InsertRequest) error
}

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgres{
		db: db,
	}
}

func (p postgres) InsertRequest(ctx context.Context, req *InsertRequest) error {
	q := `INSERT INTO request (method, path, code, "time") VALUES($1, $2, $3, $4);`

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	stmt, err := p.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		req.Method,
		req.Path,
		req.Code,
		req.Time,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		return errors.New("expected affects a row")
	}

	return nil
}
