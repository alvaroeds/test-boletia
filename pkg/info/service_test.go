package info

import (
	"context"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.Background()

type repositoryMock struct{}

func (r repositoryMock) InsertRequest(ctx context.Context, req *InsertRequest) error {
	//TODO implement me
	panic("implement me")
}

type insertQueryMock struct {
	repositoryMock
	insertQueryErr error
}

func (m insertQueryMock) InsertRequest(ctx context.Context, req *InsertRequest) error {
	return m.insertQueryErr
}

func TestService_InsertQuery(t *testing.T) {
	errorMock := errors.New("not found")
	tests := []struct {
		name           string
		insertQueryReq *InsertRequest
		insertQueryErr error
		insertErr      error
	}{
		{
			name: "success",
			insertQueryReq: &InsertRequest{
				Method: "GET",
				Path:   "/test",
				Code:   200,
				Time:   0.12,
			},
			insertQueryErr: nil,
			insertErr:      nil,
		},
		{
			name:           "fail",
			insertQueryReq: &InsertRequest{},
			insertQueryErr: errorMock,
			insertErr:      errorMock,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := insertQueryMock{
				insertQueryErr: test.insertQueryErr,
			}

			s := service{
				repo: repo,
			}
			err := s.InsertRequest(ctx, test.insertQueryReq)
			assert.Equal(t, test.insertErr, err)
		})
	}
}

func TestService_newServicee(t *testing.T) {
	NewService(&sql.DB{})
}
