package currency

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"
)

type Repository interface {
	GetAll(ctx context.Context) ([]GetAllCurrency, error)
	GetByID(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error)
	Insert(ctx context.Context, req *InsertCurrency) error
}

const timeout = 5 * time.Second

type postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgres{
		db: db,
	}
}

func (p postgres) GetAll(ctx context.Context) ([]GetAllCurrency, error) {
	q := `select DISTINCT ON (code) code,"value", created_at from currency order by code, created_at desc;`

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res := make([]GetAllCurrency, 0)
	rows, err := p.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c GetAllCurrency
		err = rows.Scan(
			&c.Code,
			&c.Value,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return res, nil
}

func (p postgres) GetByID(ctx context.Context, f CurrencyFilterRequest) ([]GetAllCurrency, error) {
	var q strings.Builder

	q.WriteString(`SELECT code, "value", created_at FROM currency`)
	args := make([]interface{}, 0)
	q.WriteString(" WHERE")
	if f.Code != "" {
		args = append(args, f.Code)
		q.WriteString(" code = $")
		q.WriteString(strconv.Itoa(len(args)))
	}

	if !f.FEnd.IsZero() {
		if len(args) > 0 {
			q.WriteString(" AND")
		}
		args = append(args, f.FInit, f.FEnd)
		q.WriteString(" created_at BETWEEN $")
		q.WriteString(strconv.Itoa(len(args) - 1))
		q.WriteString(` AND $`)
		q.WriteString(strconv.Itoa(len(args)))
	} else {
		q.WriteString("order by created_at desc")
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	res := make([]GetAllCurrency, 0)
	query := q.String()
	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c GetAllCurrency
		err = rows.Scan(
			&c.Code,
			&c.Value,
			&c.CreatedAt,
		)
		if err != nil {
			continue
		}

		res = append(res, c)
	}

	return res, nil
}

func (p postgres) Insert(ctx context.Context, req *InsertCurrency) error {
	q := `INSERT INTO currency (code, "value") VALUES($1, $2);`

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	stmt, err := p.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		req.Code,
		req.Value,
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
